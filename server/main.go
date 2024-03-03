package main

import (
	"fmt"

	"github.com/RedCuckoo/merkle-tree-verifier/server/src/config"
	"github.com/RedCuckoo/merkle-tree-verifier/server/src/manager"
	"github.com/RedCuckoo/merkle-tree-verifier/server/src/server"
)

func main() {
	cfg := config.NewConfig()
	storageManager := manager.NewManager()
	// Create gRPC server
	merkleTreeServer := server.NewMerkleTreeServer(cfg, storageManager)

	// Start gRPC server
	if err := merkleTreeServer.Start(); err != nil {
		panic(fmt.Sprintf("Failed to start gRPC server: %v", err))
	}
}
