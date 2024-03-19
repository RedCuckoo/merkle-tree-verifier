package client

import (
	"bytes"
	"context"
	cryptorand "crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/RedCuckoo/merkle-tree-verifier/pkg/merkle_tree"
	proto "github.com/RedCuckoo/merkle-tree-verifier/proto/generated"
)

const (
	ROOT_DIR         = "./storage"
	STORAGE_DIR      = ROOT_DIR + "/" + "files"
	MERKLE_ROOT_FILE = "0081cde4ac32f0f9222218e0c29a4923a750b2cf3576a3604b4395402d6b89ed"
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
	logger := log.New(os.Stdout, "[client] ", log.LstdFlags)
	err := os.MkdirAll(STORAGE_DIR, 0o755)
	if err != nil {
		log.Fatalf("failed to initialize client, %s\n", err.Error())
	}

	root, err := readMerkleRootFromFile(ROOT_DIR)
	if err != nil {
		log.Fatalf("failed to initialize client, %s\n", err.Error())
	}

	return &Service{
		merkleTreeCreated:  len(root) != 0,
		logger:             logger,
		merkleTreeService:  merkleTreeService,
		merkleTreeRootHash: root,
	}
}

func readMerkleRootFromFile(dir string) ([]byte, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return make([]byte, 0), nil
		}
		return nil, err
	}

	if len(files) == 0 {
		return make([]byte, 0), nil
	}

	filePath := filepath.Join(dir, MERKLE_ROOT_FILE)

	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return make([]byte, 0), nil
		}
		return make([]byte, 0), fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	return content, nil
}

func writeMerkleRootToFile(dir string, root []byte) error {
	filePath := filepath.Join(dir, MERKLE_ROOT_FILE)

	err := os.WriteFile(filePath, root, 0o644)
	if err != nil {
		return fmt.Errorf("error writing file %s: %w", filePath, err)
	}

	return nil
}

func (c *Service) GenerateFiles(amount int) error {
	if c.merkleTreeCreated {
		return ErrTreeCreated
	}

	err := os.MkdirAll(STORAGE_DIR, 0o755)
	if err != nil {
		log.Fatalf("failed to create directory, %s\n", err.Error())
	}

	c.logger.Printf("generating %d files...\n", amount)

	for i := 1; i <= amount; i++ {
		filename := fmt.Sprintf("%s/file%d.txt", STORAGE_DIR, i)
		err := os.WriteFile(filename, generateRandomData(), 0o644)
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
		return ErrGeneratedNotCalled
	}

	files, fileNames, err := getFilesInRepo(STORAGE_DIR)
	if err != nil {
		return err
	}

	c.merkleTreeRootHash = new(merkle_tree.MerkleTree).Init(files).GetRoot()

	err = writeMerkleRootToFile(ROOT_DIR, c.merkleTreeRootHash)
	if err != nil {
		return err
	}

	reply, err := c.merkleTreeService.UploadFiles(
		context.Background(),
		&proto.UploadFilesRequest{Files: files, FileNames: fileNames},
	)
	if err != nil {
		return err
	}

	if !bytes.Equal(reply.MerkleTreeRoot, c.merkleTreeRootHash) {
		return ErrServerRootMismatch
	}

	os.RemoveAll(STORAGE_DIR)

	log.Println("unloaded and deleted local copies successfully")

	return nil
}

func (c *Service) Download(index uint64) error {
	if index == 0 {
		return errors.New("invalid index")
	}

	index--

	reply, err := c.merkleTreeService.DownloadFile(
		context.Background(),
		&proto.DownloadFileRequest{FileIndex: index},
	)
	if err != nil {
		return err
	}

	ok, err := new(merkle_tree.MerkleTreeVerifier).ValidateFileByProof(
		reply.MerkleProof.GetProvenData(),
		new(merkle_tree.MerkleProof).UnmarshalProto(reply.MerkleProof),
		c.merkleTreeRootHash,
	)
	if err != nil {
		return err
	}

	if ok {
		err := os.MkdirAll(STORAGE_DIR, 0o755)
		if err != nil {
			log.Fatalf("failed to create directory, %s\n", err.Error())
		}

		filePath := filepath.Join(STORAGE_DIR, reply.GetFileName())

		err = os.WriteFile(filePath, reply.GetMerkleProof().GetProvenData(), 0o644)
		if err != nil {
			return fmt.Errorf("error writing to file %s: %w", filePath, err)
		}

		c.logger.Println("downloaded and verified successfully")
	} else {
		c.logger.Println("downloaded but failed to verify, FILE IS CORRUPT")
	}

	return nil
}

func (c *Service) ListLocal() error {
	files, err := getStoredFiles(STORAGE_DIR)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "Listing local files:")
	for _, file := range files {
		fmt.Fprintf(os.Stdout, "* %s\n", file.Name())
	}

	return nil
}

func (c *Service) ListRemote() error {
	remote, err := c.merkleTreeService.ListRemote(context.Background(), &proto.ListRemoteRequest{})
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "Listing remote files:")
	for i, file := range remote.GetFileNames() {
		fmt.Fprintf(os.Stdout, "%d. %s\n", i+1, file)
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
	} else {
		c.logger.Println("reset server successfully")
	}

	if err = os.RemoveAll(ROOT_DIR); err == nil {
		c.logger.Println("reset client successfully")
		return nil
	} else {
		c.logger.Println("failed to reset client")
		return err
	}
}

func getStoredFiles(dir string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return make([]os.DirEntry, 0), nil
		}
		return nil, err
	}
	return files, nil
}

func getFilesInRepo(repoPath string) ([][]byte, []string, error) {
	var files [][]byte
	var fileNames []string
	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			files = append(files, data)
			fileNames = append(fileNames, info.Name())
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return files, fileNames, nil
}

func generateRandomData() []byte {
	// Generate random length between 10 and 100 bytes
	length := rand.Intn(91) + 10
	data := make([]byte, length)

	if _, err := cryptorand.Read(data); err != nil {
		return []byte{}
	}
	return data
}
