package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	proto "github.com/RedCuckoo/merkle-tree-verifier/proto/generated"
	"github.com/RedCuckoo/merkle-tree-verifier/server/src/config"
	"github.com/RedCuckoo/merkle-tree-verifier/server/src/manager"
	"google.golang.org/grpc"
)

type MerkleTreeServer struct {
	config  *config.Config
	server  *grpc.Server
	manager *manager.Manager
	logger  *log.Logger
	proto.MerkleTreeServerServer
}

func NewMerkleTreeServer(config *config.Config, manager *manager.Manager) *MerkleTreeServer {
	logger := log.New(os.Stdout, "[client] ", log.LstdFlags)
	return &MerkleTreeServer{
		config:  config,
		manager: manager,
		logger:  logger,
	}
}

func (s *MerkleTreeServer) Start() error {
	lis, err := net.Listen("tcp", s.config.GRPCAddress)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.server = grpc.NewServer()

	proto.RegisterMerkleTreeServerServer(s.server, s)

	s.logger.Printf("Starting gRPC server on %s...\n", s.config.GRPCAddress)
	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
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
	return s.manager.UploadFiles(ctx, request)
}

func (s *MerkleTreeServer) DownloadFile(
	ctx context.Context,
	request *proto.DownloadFileRequest,
) (*proto.DownloadFileReply, error) {
	return s.manager.DownloadFile(ctx, request)
}

func (s *MerkleTreeServer) ListRemote(
	ctx context.Context,
	request *proto.ListRemoteRequest,
) (*proto.ListRemoteReply, error) {
	return s.manager.ListRemote(ctx, request)
}

func (s *MerkleTreeServer) Reset(
	ctx context.Context,
	request *proto.ResetRequest,
) (*proto.ResetReply, error) {
	return s.manager.Reset(ctx, request)
}
