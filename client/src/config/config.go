package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ServerGRPCAddress string `mapstructure:"server-grpc-address"`
}

func NewConfig() *Config {
	configFilePath := os.Getenv("CONFIG_FILE")
	if configFilePath == "" {
		panic("CONFIG_FILE_PATH environment variable is not set")
	}
	namespace := "merkle-tree-client"

	v := viper.New()
	v.AddConfigPath("./configs/")
	v.SetConfigName("config.client")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Error reading config file: %s\n", err))
	}

	v = v.Sub(namespace)
	if v == nil {
		panic("config has no values under namespace")
	}

	var config Config

	err = v.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprintf("Error parsing config file: %s\n", err))
	}

	return &config
}
