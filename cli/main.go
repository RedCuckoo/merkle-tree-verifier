package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/RedCuckoo/merkle-tree-verifier/client/src/client"
	"github.com/RedCuckoo/merkle-tree-verifier/client/src/config"
	proto "github.com/RedCuckoo/merkle-tree-verifier/proto/generated"
	prompt "github.com/c-bata/go-prompt"
	"go.uber.org/dig"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	container := dig.New()
	must(container.Provide(func() *dig.Container { return container }))
	must(container.Provide(config.NewConfig))
	must(container.Provide(func(cfg *config.Config) proto.MerkleTreeServerClient {
		merkleTreeServerClient, err := grpc.Dial(
			cfg.ServerGRPCAddress,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(419430400)),
		)
		if err != nil {
			log.Fatalf(
				"failed to connect to server by address %s: %s",
				cfg.ServerGRPCAddress,
				err,
			)
		}
		return proto.NewMerkleTreeServerClient(merkleTreeServerClient)
	}))
	must(container.Provide(client.NewClientService))

	must(container.Invoke(func(
		cfg *config.Config,
		clientService *client.Service,
	) {
		executor := NewCommandExecutor(clientService)

		for {
			p := prompt.Input(
				"merkle-tree-verifier> ",
				completer,
				prompt.OptionTitle("merkle-tree-verifier CLI"),
				prompt.OptionPrefixBackgroundColor(prompt.Purple),
			)

			executor.ExecuteCommand(p)

		}
	}))

	<-ctx.Done()
	time.Sleep(1 * time.Second)
}

func must(err error) {
	if err != nil {
		log.Fatalf("failed to initialize DI: %s", err)
	}
}
