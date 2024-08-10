package repositories

import (
	"diving-log-book-service/internal/pkg/apihelper"
	"diving-log-book-service/internal/types"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var client *s3.S3

//go:generate mockgen -source=storage.go -destination=../../test/repositories/mock/storage.go

type StorageInterface interface {
	GetUrl(*types.GetUrl) (string, error)
	Upload(*types.UploadPayload) (*types.UploadedFile, error)
	Delete(*types.DeleteObject) error
}

type StorageRepository struct{}

func NewStorageRepository(config *aws.Config) StorageInterface {
	if client == nil {
		client = connection(config)
	}

	return &StorageRepository{}
}

func (s StorageRepository) GetUrl(obj *types.GetUrl) (string, error) {
	exist, err := checkObjectExists(obj.Bucket, obj.Key)
	if err != nil || !exist {
		return "", err
	}

	input := &s3.GetObjectInput{
		Bucket: &obj.Bucket,
		Key:    &obj.Key,
	}

	req, _ := client.GetObjectRequest(input)
	url, err := req.Presign(45 * time.Minute)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s StorageRepository) Upload(obj *types.UploadPayload) (*types.UploadedFile, error) {
	input := &s3.PutObjectInput{
		Bucket: &obj.Bucket,
		Key:    &obj.Name,
		Body:   obj.Body,
	}

	_, err := client.PutObject(input)
	if err != nil {
		return nil, err
	}

	res := &types.UploadedFile{
		Bucket: obj.Bucket,
		Key:    obj.Name,
	}

	return res, nil
}

func (s StorageRepository) Delete(obj *types.DeleteObject) error {
	input := &s3.DeleteObjectInput{
		Bucket: &obj.Bucket,
		Key:    &obj.Key,
	}

	_, err := client.DeleteObject(input)
	if err != nil {
		return err
	}

	return nil
}

func checkObjectExists(bucket string, key string) (bool, error) {
	_, err := client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return false, apihelper.S3Error(err)
	}

	return true, nil
}

func connection(config *aws.Config) *s3.S3 {
	sess, err := session.NewSession(config)

	if err != nil {
		log.Fatal("Failed to create session", err)
	}

	return s3.New(sess)
}
