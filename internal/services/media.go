package services

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
)

type MediaService struct {
	mediaRepo repositories.MediaInterface
}

func NewMediaService(repo repositories.MediaInterface) *MediaService {
	return &MediaService{
		mediaRepo: repo,
	}
}

func (m MediaService) Create(payload *types.CreateMediaPayload) (*models.Media, error) {
	media, err := m.mediaRepo.Create(payload)
	if err != nil {
		return nil, err
	}

	return media, nil
}
