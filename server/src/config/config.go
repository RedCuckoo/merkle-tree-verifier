package config

import (
	"github.com/RedCuckoo/merkle-tree-verifier/utils"
)

type Config struct {
	GRPCAddress string `mapstructure:"grpc-address"`
}

func NewConfig() *Config {
	config := Config{}

	utils.Load("merkle-tree-server", &config)

	return &config
}
