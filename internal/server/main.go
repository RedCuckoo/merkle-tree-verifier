package main

import (
	"fmt"

	"github.com/RedCuckoo/merkle-tree-verifier/internal/server/config"
	"github.com/RedCuckoo/merkle-tree-verifier/internal/server/manager"
	"github.com/RedCuckoo/merkle-tree-verifier/internal/server/server"
)

func main() {
	cfg := config.NewConfig()
	storageManager := manager.NewManager()
	merkleTreeServer := server.NewMerkleTreeServer(cfg, storageManager)

	if err := merkleTreeServer.Start(); err != nil {
		panic(fmt.Sprintf("Failed to start gRPC server: %v", err))
	}
}
