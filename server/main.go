package main

import (
	"fmt"

	"github.com/RedCuckoo/merkle-tree-verifier/server/src/config"
	"github.com/RedCuckoo/merkle-tree-verifier/server/src/server"
)

func main() {
	cfg := config.NewConfig()

	// Create gRPC server
	merkleTreeServer := server.NewMerkleTreeServer(cfg)

	// Start gRPC server
	if err := merkleTreeServer.Start(); err != nil {
		panic(fmt.Sprintf("Failed to start gRPC server: %v", err))
	}
}
