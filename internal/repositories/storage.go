package repositories

import (
	"bytes"
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
	InitMultipart(multipart *types.InitMultipartPayload) (*s3.CreateMultipartUploadOutput, error)
	UploadPart(input *types.UploadPartPayload) (*s3.UploadPartOutput, error)
	CompleteMultipart(payload *types.CompleteMultipartPayload) (*s3.CompleteMultipartUploadOutput, error)
	AbortMultipart(payload *types.AbortMultipartPayload) error
}

type StorageRepository struct {
}

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
		Body:   obj.Body,
		Key:    &obj.Key,
	}

	_, err := client.PutObject(input)
	if err != nil {
		return nil, err
	}

	res := &types.UploadedFile{
		Bucket: obj.Bucket,
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

func (s *StorageRepository) InitMultipart(multipart *types.InitMultipartPayload) (*s3.CreateMultipartUploadOutput, error) {
	input := &s3.CreateMultipartUploadInput{
		Bucket: aws.String(multipart.Bucket),
		Key:    aws.String(multipart.Key),
	}

	multipartUpload, err := client.CreateMultipartUpload(input)
	if err != nil {
		return nil, err
	}

	return multipartUpload, nil
}

func (s StorageRepository) UploadPart(input *types.UploadPartPayload) (*s3.UploadPartOutput, error) {
	part, err := client.UploadPart(&s3.UploadPartInput{
		Bucket:     aws.String(input.Bucket),
		Key:        aws.String(input.Key),
		PartNumber: aws.Int64(int64(input.Part)),
		UploadId:   aws.String(input.UploadID),
		Body:       bytes.NewReader(input.Buffer),
	})
	if err != nil {
		return nil, err
	}

	return part, nil
}

func (s StorageRepository) AbortMultipart(payload *types.AbortMultipartPayload) error {
	input := &s3.AbortMultipartUploadInput{
		Bucket:   &payload.Bucket,
		Key:      &payload.Key,
		UploadId: &payload.UploadID,
	}
	_, err := client.AbortMultipartUpload(input)
	if err != nil {
		return err
	}

	return nil
}

func (s StorageRepository) CompleteMultipart(payload *types.CompleteMultipartPayload) (*s3.CompleteMultipartUploadOutput, error) {
	res, err := client.CompleteMultipartUpload(&s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(payload.Bucket),
		Key:      aws.String(payload.Key),
		UploadId: aws.String(payload.UploadID),
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: payload.MultipartUpload,
		},
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func connection(config *aws.Config) *s3.S3 {
	sess, err := session.NewSession(config)

	if err != nil {
		log.Fatal("Failed to create session", err)
	}

	return s3.New(sess)
}
