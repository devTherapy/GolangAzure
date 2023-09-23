package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config Configuration

type Configuration struct {
	ContainerName    string
	ConnectionString string
	StorageName      string
	AccountKey       string
}

func SetupConfig(configPath string) {
	var config Configuration
	viper.SetConfigFile(configPath)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	Config = config
}

func GetConfig() Configuration {
	return Config
}
