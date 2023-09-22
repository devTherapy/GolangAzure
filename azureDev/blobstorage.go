package azureDev

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type FileType int32 //enum
const (
	_             = iota
	Jpeg FileType = 1
	Txt  FileType = 2
)

type FileProperties struct {
	Filetype FileType
	Filepath string
	Filename string
}

const url = "https://sayostorage.blob.core.windows.net/"

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		log.Fatal(err)
	}
}

func UploadFile(fileProperties FileProperties) {
	containerName := "sayocontainer"
	ctx := context.Background()
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	file, err := os.Open(fileProperties.Filepath)
	handleError(err)
	defer file.Close()

	_, err = client.UploadStream(ctx, containerName, fileProperties.Filename, file, &azblob.UploadStreamOptions{})
	handleError(err)
}

func DownloadFile(fileProperties FileProperties) {
	containerName := "sayocontainer"
	ctx := context.Background()
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	file, err := os.Create(fileProperties.Filename)
	handleError(err)
	defer file.Close()

	downloadResponse, err := client.DownloadStream(ctx, containerName, fileProperties.Filename, &azblob.DownloadStreamOptions{})
	handleError(err)
	downloadResponse.Body.Close()
	_, err = io.Copy(file, downloadResponse.Body)
	handleError(err)
}

func DeleteFile(fileProperties FileProperties) {
	containerName := "sayocontainer"
	ctx := context.Background()
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	_, err = client.DeleteBlob(ctx, containerName, fileProperties.Filename, nil)
	handleError(err)
}

func ListFiles() {
	containerName := "sayocontainer"
	ctx := context.Background()
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	p := client.NewListBlobsFlatPager(containerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Versions: true, Snapshots: true},
	})

	for p.More() {
		resp, err := p.NextPage(ctx)
		handleError(err)
		for _, blob := range resp.Segment.BlobItems {
			log.Printf("Blob name: %s\n", *blob.Name)
		}
	}
}

//write a function to read a file from disc and upload it to azure blob storage
