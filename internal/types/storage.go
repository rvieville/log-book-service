package types

import "mime/multipart"

type UploadPayload struct {
	Name   string         `validate:"required"`
	Body   multipart.File `validate:"required"`
	Bucket string         `validate:"required"`
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
	Bucket string
	Key    string
}
