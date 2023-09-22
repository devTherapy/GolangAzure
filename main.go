package main

import (
	"golang-azure/azureDev"
)

func main() {

	f := azureDev.FileProperties{
		Filetype: azureDev.Jpeg,
		Filepath: "./azureDev/upload.txt",
	}
	azureDev.UploadFile(f)
}
