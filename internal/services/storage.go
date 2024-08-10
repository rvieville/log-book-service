package services

import (
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
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

func (s StorageService) Upaload(file *types.UploadPayload) (*types.UploadedFile, error) {
	obj, err := s.storageRepo.Upload(file)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (s StorageService) Delete(obj *types.DeleteObject) error {
	err := s.storageRepo.Delete(obj)
	if err != nil {
		return err
	}

	return nil
}
