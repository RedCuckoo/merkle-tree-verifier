package config

import (
	"github.com/RedCuckoo/merkle-tree-verifier/internal/utils"
)

type Config struct {
	ServerGRPCAddress string `mapstructure:"server-grpc-address"`
}

func NewConfig() *Config {
	config := Config{}

	utils.Load("merkle-tree-client", &config)

	return &config
}
