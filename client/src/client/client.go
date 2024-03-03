package client

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/RedCuckoo/merkle-tree-verifier/merkle_tree"
	proto "github.com/RedCuckoo/merkle-tree-verifier/proto/generated"
)

const (
	STORAGE_DIR = "./storage"
)

var errorTreeCreated = errors.New(
	"merkle tree already created or \"generate\" has been already called, run \"reset\" to start over",
)
var errorGeneratedNotCalled = errors.New("no files, call \"generate\" first")

var errorServerRootMismatch = errors.New(
	"merkle root tree returned from the server differs from calculated",
)

type Service struct {
	merkleTreeCreated  bool
	logger             *log.Logger
	merkleTreeRootHash []byte
	merkleTreeService  proto.MerkleTreeServerClient
}

func NewClientService(
	merkleTreeService proto.MerkleTreeServerClient,
) *Service {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	err := os.MkdirAll(STORAGE_DIR, 0o755)
	if err != nil {
		log.Fatalf("failed to initialize client, %s\n", err.Error())
	}

	files, err := os.ReadDir(STORAGE_DIR)
	if err != nil {
		log.Fatalf("failed to initialize client, %s\n", err.Error())
	}

	return &Service{
		merkleTreeCreated: len(files) != 0,
		logger:            logger,
		merkleTreeService: merkleTreeService,
	}
}

func (c *Service) GenerateFiles(amount int) error {
	if c.merkleTreeCreated {
		return errorTreeCreated
	}

	c.logger.Printf("creating storage directory\n")

	err := os.MkdirAll(STORAGE_DIR, 0o755)
	if err != nil {
		log.Fatalf("failed to create directory, %s\n", err.Error())
	}

	c.logger.Printf("directory created successfully\n")
	c.logger.Printf("generating %d files...\n", amount)

	for i := 1; i <= amount; i++ {
		filename := fmt.Sprintf("%s/file%d.txt", STORAGE_DIR, i)
		err := ioutil.WriteFile(filename, generateRandomData(), 0o644)
		if err != nil {
			return err
		}
	}

	c.merkleTreeCreated = true
	c.logger.Printf("files generated successfully\n")

	return nil
}

func (c *Service) Unload() error {
	if !c.merkleTreeCreated {
		return errorGeneratedNotCalled
	}

	files, err := getFilesInRepo(STORAGE_DIR)
	if err != nil {
		return err
	}

	c.merkleTreeRootHash = new(merkle_tree.MerkleTree).Init(files).GetRoot()

	log.Println(hex.EncodeToString(c.merkleTreeRootHash))

	reply, err := c.merkleTreeService.UploadFiles(
		context.Background(),
		&proto.UploadFilesRequest{Files: files},
	)
	if err != nil {
		return err
	}

	if !bytes.Equal(reply.MerkleTreeRoot, c.merkleTreeRootHash) {
		return errorServerRootMismatch
	}

	os.RemoveAll(STORAGE_DIR)

	return nil
}

func (c *Service) Download(index uint64) error {
	reply, err := c.merkleTreeService.DownloadFile(
		context.Background(),
		&proto.DownloadFileRequest{FileIndex: index},
	)
	if err != nil {
		return err
	}

	ok, err := new(merkle_tree.MerkleTreeVerifier).ValidateFileByProof(
		reply.File,
		new(merkle_tree.MerkleProof).UnmarshalProto(reply.MerkleProof),
		c.merkleTreeRootHash,
	)
	if err != nil {
		return err
	}

	if ok {
		c.logger.Println("downloaded and verified successfully")
	} else {
		c.logger.Println("downloaded but failed to verify")
	}

	return nil
}

func (c *Service) ListLocal() error {
	files, err := getStoredFiles()
	if err != nil {
		return err
	}

	fmt.Println("Listing local files:")
	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file.Name())
	}

	return nil
}

func (c *Service) ListRemote() error {
	remote, err := c.merkleTreeService.ListRemote(context.Background(), &proto.ListRemoteRequest{})
	if err != nil {
		return err
	}

	fmt.Println("Listing remote files:")
	for i, file := range remote.GetFileNames() {
		fmt.Printf("%d. %s\n", i+1, file)
	}

	return nil
}

func (c *Service) Reset() error {
	c.merkleTreeCreated = false
	c.merkleTreeRootHash = nil
	reset, err := c.merkleTreeService.Reset(context.Background(), &proto.ResetRequest{})
	if err != nil {
		return err
	}

	if !reset.Successful {
		c.logger.Println("failed to reset server")
	}

	return os.RemoveAll(STORAGE_DIR)
}

func getStoredFiles() ([]os.DirEntry, error) {
	return os.ReadDir(STORAGE_DIR)
}

func getFilesInRepo(repoPath string) ([][]byte, error) {
	var files [][]byte

	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			files = append(files, data)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func generateRandomData() []byte {
	// Generate random length between 10 and 100 bytes
	length := rand.Intn(91) + 10
	data := make([]byte, length)
	rand.Read(data)
	return data
}
