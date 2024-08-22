package repositories_test

import (
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
	"diving-log-book-service/test/utils"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
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
		Body:   file,
		Key:    "fish.txt",
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
		Key:    "fish.txt",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	storage := getStorageRepo()
	input := &types.DeleteObject{
		Bucket: "dive",
		Key:    "fish.txt",
	}

	err := storage.Delete(input)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitMultipart(t *testing.T) {
	storage := getStorageRepo()
	input := &types.InitMultipartPayload{
		Bucket: "dive",
		Key:    "fish.txt",
	}

	res, err := storage.InitMultipart(input)
	if err != nil {
		t.Fatal(err)
	}

	utils.UploadID = res.UploadId
}

func TestUploadPart(t *testing.T) {
	utils.Parts = []*s3.CompletedPart{}
	storage := getStorageRepo()
	file, err := os.Open("./files/fish.txt")
	if err != nil {
		t.Fatal(err)
	}
	buff := make([]byte, 5*1024*1024)
	part := int64(1)

	file.Read(buff)

	input := &types.UploadPartPayload{
		Bucket:   "dive",
		Key:      "fish.txt",
		Part:     part,
		UploadID: *utils.UploadID,
		Buffer:   buff,
	}
	res, err := storage.UploadPart(input)
	if err != nil {
		t.Fatal(err)
	}

	utils.Parts = append(utils.Parts, &s3.CompletedPart{
		ETag:       res.ETag,
		PartNumber: &part,
	})
}

func TestCompleteMultipart(t *testing.T) {
	storage := getStorageRepo()

	input := &types.CompleteMultipartPayload{
		Bucket:          "dive",
		Key:             "fish.txt",
		UploadID:        *utils.UploadID,
		MultipartUpload: utils.Parts,
	}
	_, err := storage.CompleteMultipart(input)
	if err != nil {
		t.Fatal(err)
	}

	utils.Parts = []*s3.CompletedPart{}
	utils.UploadID = nil
}
