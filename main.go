package main

import (
	"golang-azure/azureDev"
	"golang-azure/azureDev/config"
	"log"

	"github.com/spf13/viper"
)

func main() {

	config.SetupConfig("./azureDev/config/config.json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	f := azureDev.FileProperties{
		Filename: "upload.txt",
		Filetype: "txt",
		Filepath: "./azureDev/upload.txt",
	}
	azureDev.DeleteFile(f)
}
