package services

import (
	"diving-log-book-service/internal/models"
	gormHelper "diving-log-book-service/internal/pkg/gorm"
	"diving-log-book-service/internal/repositories"
)

type IslandService struct {
	islandRepo repositories.IslandInterface
}

func NewIslandService(islandRepo repositories.IslandInterface) *IslandService {
	return &IslandService{
		islandRepo,
	}
}

func (i IslandService) ReadAll(filters []gormHelper.Filter) ([]models.Island, error) {
	islands, err := i.islandRepo.ReadAll(filters)
	if err != nil {
		return nil, err
	}

	return islands, nil
}
