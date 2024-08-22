package types

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3"
)

type UploadPayload struct {
	Bucket string
	Body   multipart.File
	Key    string
}

type CompleteMultipartPayload struct {
	Bucket          string
	Key             string
	UploadID        string
	MultipartUpload []*s3.CompletedPart
}

type AbortMultipartPayload struct {
	Bucket   string
	Key      string
	UploadID string
}

type UploadPartPayload struct {
	Bucket   string
	Key      string
	Part     int64
	UploadID string
	Buffer   []byte
}

type GetUrl struct {
	Bucket string `validate:"required"`
	Key    string `validate:"required"`
}

type DeleteObject struct {
	Bucket string `validate:"required"`
	Key    string `validate:"required"`
}

type UploadedFile struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type InitMultipartPayload struct {
	Bucket string
	Key    string
}
