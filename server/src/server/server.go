package server

import (
	"context"
	"fmt"
	"github.com/RedCuckoo/merkle-tree-verifier/server/src/manager"
	"net"

	proto "github.com/RedCuckoo/merkle-tree-verifier/proto/generated"
	"github.com/RedCuckoo/merkle-tree-verifier/server/src/config"

	"google.golang.org/grpc"
)

type MerkleTreeServer struct {
	config  *config.Config
	server  *grpc.Server
	manager *manager.Manager
	proto.MerkleTreeServerServer
}

func NewMerkleTreeServer(config *config.Config, manager *manager.Manager) *MerkleTreeServer {
	return &MerkleTreeServer{
		config:  config,
		manager: manager,
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

func (s *MerkleTreeServer) UploadFiles(
	ctx context.Context,
	request *proto.UploadFilesRequest,
) (*proto.UploadFilesReply, error) {
	return nil, nil
}
func (s *MerkleTreeServer) DownloadFile(
	ctx context.Context,
	request *proto.DownloadFileRequest,
) (*proto.DownloadFileReply, error) {
	return nil, nil
}
func (s *MerkleTreeServer) ListRemote(
	ctx context.Context,
	request *proto.ListRemoteRequest,
) (*proto.ListRemoteReply, error) {
	return nil, nil
}
func (s *MerkleTreeServer) Reset(
	ctx context.Context,
	request *proto.ResetRequest,
) (*proto.ResetReply, error) {
	return nil, nil
}
