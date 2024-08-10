package repositories_test

import (
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func getStorageRepo() repositories.StorageInterface {
	config := &aws.Config{
		Region:   aws.String("ap-southeast-1"),
		Endpoint: aws.String("localhost:9000"),
		Credentials: credentials.NewStaticCredentials(
			"WGo3v9EPOOGeDEn1GOs1",
			"MStlaXHGnTLj8b5KrPu2mc4lijobbRt7O1nnpVfu",
			"", // a token will be created when the session it's used.
		),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL:       aws.Bool(true),
	}

	return repositories.NewStorageRepository(config)
}

func TestUpload(t *testing.T) {
	storageRepo := getStorageRepo()

	file, err := os.Open("./files/fish.txt")
	if err != nil {
		t.Fatal(err)
	}

	input := &types.UploadPayload{
		Bucket: "dive",
		Name:   "testUpload.txt",
		Body:   file,
	}

	_, err = storageRepo.Upload(input)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUrl(t *testing.T) {
	storage := getStorageRepo()

	_, err := storage.GetUrl(&types.GetUrl{
		Bucket: "dive",
		Key:    "testUpload.txt",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	storage := getStorageRepo()
	input := &types.DeleteObject{
		Bucket: "dive",
		Key:    "testUpload.txt",
	}

	err := storage.Delete(input)
	if err != nil {
		t.Fatal(err)
	}
}
