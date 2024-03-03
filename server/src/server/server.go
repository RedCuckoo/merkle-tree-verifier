package server

import (
	"fmt"
	"net"

	proto "github.com/RedCuckoo/merkle-tree-verifier/proto/generated"
	"github.com/RedCuckoo/merkle-tree-verifier/server/src/config"

	"google.golang.org/grpc"
)

type MerkleTreeServer struct {
	config *config.Config
	server *grpc.Server
	proto.MerkleTreeServerServer
}

func NewMerkleTreeServer(config *config.Config) *MerkleTreeServer {
	return &MerkleTreeServer{
		config: config,
	}
}

func (s *MerkleTreeServer) Start() error {
	lis, err := net.Listen("tcp", s.config.GRPCAddress)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer()

	proto.RegisterMerkleTreeServerServer(s.server, s)

	fmt.Printf("Starting gRPC server on %s...\n", s.config.GRPCAddress)
	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (s *MerkleTreeServer) Stop() {
	s.server.GracefulStop()
}
