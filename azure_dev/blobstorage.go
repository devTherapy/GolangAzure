package azure_dev

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type fileType int32 //enum
const (
	_             = iota
	jpeg fileType = 1
	txt  fileType = 2
)

type fileProperties struct {
	filetype fileType
	filepath string
	fileName string
}

const url = "https://sayostorage.blob.core.windows.net/"

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func UploadFile(fileProperties fileProperties) {
	containerName := "sayocontainer"
	ctx := context.Background()
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	file, err := os.Open(fileProperties.filepath)
	handleError(err)
	defer file.Close()

	_, err = client.UploadStream(ctx, containerName, fileProperties.fileName, file, &azblob.UploadStreamOptions{})
	handleError(err)
}

func DownloadFile(fileProperties fileProperties) {
	containerName := "sayocontainer"
	ctx := context.Background()
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	file, err := os.Create(fileProperties.fileName)
	handleError(err)
	defer file.Close()

	downloadResponse, err := client.Download(ctx, containerName, fileProperties.fileName, &azblob.DownloadStreamOptions{})
	handleError(err)

	fileBody := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})
	_, err = file.Write(fileBody)
	handleError(err)
}

//write a function to read a file from disc and upload it to azure blob storage
