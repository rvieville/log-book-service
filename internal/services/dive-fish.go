package services

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
)

type DiveFishService struct {
	repository repositories.DiveFishInterface
}

func NewDiveFishService(repository repositories.DiveFishInterface) *DiveFishService {
	return &DiveFishService{
		repository,
	}
}

func (fm DiveFishService) Create(payload *types.CreateFishPayload) (*models.DiveFish, error) {
	return fm.repository.Create(payload)
}
