package manager

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/RedCuckoo/merkle-tree-verifier/merkle_tree"
	proto "github.com/RedCuckoo/merkle-tree-verifier/proto/generated"
)

const STORAGE_DIR = "./server_storage"

type Manager struct {
	merkleTree *merkle_tree.MerkleTree
	fileNames  map[uint64]string
	filesMutex *sync.RWMutex
}

func NewManager() *Manager {
	manager := &Manager{
		filesMutex: new(sync.RWMutex),
	}

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			manager.syncManager()
		}
	}()

	return manager.syncManager()
}

func (m *Manager) syncManager() *Manager {
	m.filesMutex.Lock()
	defer m.filesMutex.Unlock()

	if err := os.MkdirAll(STORAGE_DIR, 0o755); err != nil {
		panic(fmt.Sprintf("failed to create dir: %v", err))
	}

	files, err := os.ReadDir(STORAGE_DIR)
	if err != nil {
		panic(fmt.Sprintf("failed to create dir: %v", err))
	}

	if len(files) == 0 {
		m.merkleTree = nil
		m.fileNames = nil
	} else {
		fileNames := make(map[uint64]string)
		fileContent := make([][]byte, len(files))
		for i, file := range files {
			fileNames[uint64(i)] = file.Name()

			content, err := os.ReadFile(filepath.Join(STORAGE_DIR, file.Name()))
			if err != nil {
				panic(fmt.Sprintf("failed to read file %s: %v", file.Name(), err))
			}
			fileContent[i] = content
		}

		merkleTree := new(merkle_tree.MerkleTree).Init(fileContent)

		m.merkleTree = merkleTree
		m.fileNames = fileNames
	}
	return m
}

func (m *Manager) UploadFiles(
	_ context.Context,
	request *proto.UploadFilesRequest,
) (*proto.UploadFilesReply, error) {
	m.filesMutex.Lock()
	defer m.filesMutex.Unlock()

	if m.merkleTree != nil || m.fileNames != nil {
		return nil, fmt.Errorf("merkle tree is initialized, reset first")
	}

	m.fileNames = make(map[uint64]string)
	if err := os.MkdirAll(STORAGE_DIR, 0o755); err != nil {
		return nil, err
	}

	if len(request.GetFiles()) != len(request.GetFileNames()) {
		return nil, fmt.Errorf("number of data slices does not match number of filenames")
	}

	for i, bytes := range request.GetFiles() {
		fileName := request.GetFileNames()[i]
		filePath := filepath.Join(STORAGE_DIR, fileName)

		if err := os.WriteFile(filePath, bytes, 0o644); err != nil {
			return nil, fmt.Errorf("error writing file %s: %w", fileName, err)
		}

		m.fileNames[uint64(i)] = fileName
	}

	m.merkleTree = new(merkle_tree.MerkleTree).Init(request.GetFiles())

	root := m.merkleTree.GetRoot()

	return &proto.UploadFilesReply{MerkleTreeRoot: root}, nil
}

func (m *Manager) DownloadFile(
	_ context.Context,
	request *proto.DownloadFileRequest,
) (*proto.DownloadFileReply, error) {
	m.filesMutex.RLock()
	defer m.filesMutex.RUnlock()

	if m.merkleTree == nil {
		return nil, fmt.Errorf("merkle tree wasn't initialized, unload files first")
	}

	if m.fileNames == nil {
		return nil, fmt.Errorf("storage is not initialized, reset and start over")
	}

	if request.GetFileIndex() >= uint64(len(m.fileNames)) {
		return nil, fmt.Errorf("file index out of range")
	}

	fileName := m.fileNames[request.GetFileIndex()]

	proof := m.merkleTree.GetProof(request.GetFileIndex())

	return &proto.DownloadFileReply{
		FileName:    fileName,
		MerkleProof: proof.MarshalProto(),
	}, nil
}

func (m *Manager) ListRemote(
	_ context.Context,
	_ *proto.ListRemoteRequest,
) (*proto.ListRemoteReply, error) {
	m.filesMutex.RLock()
	defer m.filesMutex.RUnlock()

	files, err := os.ReadDir(STORAGE_DIR)
	if err != nil {
		if os.IsNotExist(err) {
			return &proto.ListRemoteReply{FileNames: []string{}}, nil
		}
		return nil, err
	}

	fileNames := make([]string, len(files))
	for i, file := range files {
		fileNames[i] = file.Name()
	}

	return &proto.ListRemoteReply{FileNames: fileNames}, nil
}

func (m *Manager) Reset(
	_ context.Context,
	_ *proto.ResetRequest,
) (*proto.ResetReply, error) {
	m.filesMutex.Lock()
	defer m.filesMutex.Unlock()

	m.merkleTree = nil
	m.fileNames = nil

	if err := os.RemoveAll(STORAGE_DIR); err != nil {
		return &proto.ResetReply{Successful: false}, err
	}

	return &proto.ResetReply{Successful: true}, nil
}
