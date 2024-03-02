package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

const (
	STORAGE_DIR = "./storage"
)

var errorTreeCreated = errors.New("merkle tree already created or \"generate\" has been already called, run \"reset\" to start over")

type ClientService struct {
	merkleTreeCreated bool
	logger            *log.Logger
}

func NewClientService() *ClientService {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	err := os.MkdirAll(STORAGE_DIR, 0755)
	if err != nil {
		log.Fatalf("failed to initialize client, %s\n", err.Error())
	}

	files, err := os.ReadDir(STORAGE_DIR)

	if err != nil {
		log.Fatalf("failed to initialize client, %s\n", err.Error())
	}

	return &ClientService{
		merkleTreeCreated: len(files) != 0,
		logger:            logger,
	}
}

func (c *ClientService) GenerateFiles(amount int) error {
	if c.merkleTreeCreated {
		return errorTreeCreated
	}

	c.logger.Printf("creating storage directory\n")

	err := os.MkdirAll(STORAGE_DIR, 0755)
	if err != nil {
		log.Fatalf("failed to create directory, %s\n", err.Error())
	}

	c.logger.Printf("directory created successfully\n")
	c.logger.Printf("generating %d files...\n", amount)

	for i := 1; i <= amount; i++ {
		filename := fmt.Sprintf("%s/file%d.txt", STORAGE_DIR, i)
		err := ioutil.WriteFile(filename, generateRandomData(), 0644)
		if err != nil {
			return err
		}
	}

	c.merkleTreeCreated = true
	c.logger.Printf("files generated successfully\n")

	return nil
}

func (c *ClientService) Unload() error {
	if c.merkleTreeCreated {
		return errorTreeCreated
	}
	c.merkleTreeCreated = true

	// TODO: unload to server
	os.RemoveAll(STORAGE_DIR)

	return nil
}

func (c *ClientService) Download() error {
	// TODO: download from server
	// TODO: verify

	return nil
}

func (c *ClientService) ListLocal() error {
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

func (c *ClientService) Reset() error {
	c.merkleTreeCreated = false
	// TODO: reset server
	return os.RemoveAll(STORAGE_DIR)
}

func getStoredFiles() ([]os.DirEntry, error) {
	return os.ReadDir(STORAGE_DIR)
}

func generateRandomData() []byte {
	// Generate random length between 10 and 100 bytes
	length := rand.Intn(91) + 10
	data := make([]byte, length)
	rand.Read(data)
	return data
}
