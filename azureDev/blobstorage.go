package azureDev

import (
	"fmt"
	"golang-azure/azureDev/config"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/storage"
)

type FileType int32 //enum
const (
	_             = iota
	Jpeg FileType = 1
	Txt  FileType = 2
)

type FileProperties struct {
	Filetype string
	Filepath string
	Filename string
	Url      string
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		log.Fatal(err)
	}
}

// func UploadFile(fileProperties FileProperties) {
// 	containerName := "sayocontainer"
// 	ctx := context.Background()
// 	credential, err := azidentity.NewDefaultAzureCredential(nil)
// 	handleError(err)
// 	client, err := azblob.NewClient(url, credential, nil)
// 	handleError(err)

// 	file, err := os.Open(fileProperties.Filepath)
// 	handleError(err)
// 	defer file.Close()

// 	_, err = client.UploadStream(ctx, containerName, fileProperties.Filename, file, &azblob.UploadStreamOptions{})
// 	handleError(err)
// }

// func DownloadFile(fileProperties FileProperties) {
// 	containerName := "sayocontainer"
// 	ctx := context.Background()
// 	credential, err := azidentity.NewDefaultAzureCredential(nil)
// 	handleError(err)
// 	client, err := azblob.NewClient(url, credential, nil)
// 	handleError(err)

// 	file, err := os.Create(fileProperties.Filename)
// 	handleError(err)
// 	defer file.Close()

// 	downloadResponse, err := client.DownloadStream(ctx, containerName, fileProperties.Filename, &azblob.DownloadStreamOptions{})
// 	handleError(err)
// 	downloadResponse.Body.Close()
// 	_, err = io.Copy(file, downloadResponse.Body)
// 	handleError(err)
// }

func DeleteFile(fileProperties FileProperties) bool {

	config := config.GetConfig()
	client, err := storage.NewBasicClient(config.StorageName, config.AccountKey)
	handleError(err)

	//get reference to the container

	blobClient := client.GetBlobService()

	container := blobClient.GetContainerReference(config.ContainerName)

	if container == nil {
		log.Fatal("No container found")
	}

	blob := container.GetBlobReference(fileProperties.Filename)

	deleted, err := blob.DeleteIfExists(nil)

	handleError(err)

	return deleted
}

func UploadFile(fileProperties FileProperties) string {
	var config = config.GetConfig()

	client, err := storage.NewBasicClient(config.StorageName, config.AccountKey)
	handleError(err)

	//get reference to the container

	blobClient := client.GetBlobService()

	container := blobClient.GetContainerReference(config.ContainerName)

	if _, err := container.CreateIfNotExists(nil); err != nil {
		log.Fatal(err)
	}

	blob := container.GetBlobReference(fileProperties.Filename)
	blob.Properties.ContentType = fileProperties.Filetype
	file, err := os.Open(fileProperties.Filepath)
	handleError(err)
	defer file.Close()

	err = blob.CreateBlockBlobFromReader(file, nil)
	handleError(err)
	url := blob.GetURL()
	return url
}
