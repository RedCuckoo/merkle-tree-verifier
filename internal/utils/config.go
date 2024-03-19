package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Load(namespace string, config interface{}) {
	configFilePath := os.Getenv("CONFIG_FILE")
	if configFilePath == "" {
		panic("CONFIG_FILE environment variable is not set")
	}

	v := viper.New()
	v.AddConfigPath("./configs/")
	configFilePath = strings.Trim(configFilePath, ".yaml")
	v.SetConfigName(configFilePath)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Error reading config file: %s\n", err))
	}

	v = v.Sub(namespace)
	if v == nil {
		panic("config has no values under namespace")
	}

	err = v.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprintf("Error parsing config file: %s\n", err))
	}
}
