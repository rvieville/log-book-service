package services

import (
	"diving-log-book-service/internal/pkg/apihelper"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
)

type StorageService struct {
	storageRepo repositories.StorageInterface
}

func NewStorageService(storageRepo repositories.StorageInterface) *StorageService {
	return &StorageService{
		storageRepo,
	}
}

func (s StorageService) GetUrl(obj *types.GetUrl) (string, error) {
	url, err := s.storageRepo.GetUrl(obj)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s StorageService) Upaload(payload *types.UploadPayload) (*types.UploadedFile, error) {
	uploadID, err := s.initMultipart(payload)
	if err != nil {
		return nil, err
	}

	const partSize = 5242880 // 5 * 1024 * 1024 (5 MB)
	partNumber := int64(1)
	buffer := make([]byte, partSize)
	partETags := make([]*s3.CompletedPart, 0)

	for {
		n, err := payload.Body.Read(buffer)
		if n > 0 {
			partNb := partNumber
			eTag, err := s.uploadMultiPart(payload, partNb, uploadID, buffer)
			if err != nil {
				return nil, apihelper.InternalError(err.Error())
			}

			partETags = append(partETags, &s3.CompletedPart{
				ETag:       eTag,
				PartNumber: &partNb,
			})

			partNumber++
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			s.storageRepo.AbortMultipart(&types.AbortMultipartPayload{
				Bucket:   payload.Bucket,
				Key:      payload.Key,
				UploadID: *uploadID,
			})
			return nil, apihelper.InternalError(err.Error())
		}
	}

	uploadedFile, err := s.completeMultipart(payload, partETags, uploadID)
	if err != nil {
		s.storageRepo.AbortMultipart(&types.AbortMultipartPayload{
			Bucket:   payload.Bucket,
			Key:      payload.Key,
			UploadID: *uploadID,
		})
		return nil, apihelper.InternalError(err.Error())
	}

	res := &types.UploadedFile{
		Bucket: *uploadedFile.Bucket,
		Key:    *uploadedFile.Key,
	}

	return res, nil
}

func (s StorageService) Delete(obj *types.DeleteObject) error {
	err := s.storageRepo.Delete(obj)
	if err != nil {
		return err
	}

	return nil
}

func (s StorageService) initMultipart(payload *types.UploadPayload) (*string, error) {
	multipartObj := &types.InitMultipartPayload{
		Bucket: payload.Bucket,
		Key:    payload.Key,
	}
	multipart, err := s.storageRepo.InitMultipart(multipartObj)
	if err != nil {
		return nil, apihelper.InternalError(err.Error())
	}

	return multipart.UploadId, nil
}

func (s StorageService) uploadMultiPart(payload *types.UploadPayload, partNumber int64, uploadID *string, buffer []byte) (*string, error) {
	uploadPartPaylaod := &types.UploadPartPayload{
		Bucket:   payload.Bucket,
		Key:      payload.Key,
		Part:     partNumber,
		UploadID: *uploadID,
		Buffer:   buffer,
	}
	partResp, err := s.storageRepo.UploadPart(uploadPartPaylaod)
	if err != nil {
		s.storageRepo.AbortMultipart(&types.AbortMultipartPayload{
			Bucket:   payload.Bucket,
			Key:      payload.Key,
			UploadID: *uploadID,
		})
		return nil, apihelper.InternalError(err.Error())
	}

	return partResp.ETag, nil
}

func (s StorageService) completeMultipart(payload *types.UploadPayload, eTags []*s3.CompletedPart, uploadID *string) (*s3.CompleteMultipartUploadOutput, error) {
	completPayload := &types.CompleteMultipartPayload{
		Bucket:          payload.Bucket,
		Key:             payload.Key,
		MultipartUpload: eTags,
		UploadID:        *uploadID,
	}
	uploadedFile, err := s.storageRepo.CompleteMultipart(completPayload)
	if err != nil {
		s.storageRepo.AbortMultipart(&types.AbortMultipartPayload{
			Bucket:   payload.Bucket,
			Key:      payload.Key,
			UploadID: *uploadID,
		})
		return nil, apihelper.InternalError(err.Error())
	}

	return uploadedFile, nil
}
