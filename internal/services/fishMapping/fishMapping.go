package fishmapping

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
)

type FishMapping struct {
	repository repositories.FishMappingInterface
}

func New(repository repositories.FishMappingInterface) *FishMapping {
	return &FishMapping{
		repository,
	}
}

func (fm FishMapping) Create(payload types.CreateFishMappingPayload) (*models.FishMapping, error) {
	return fm.repository.Create(payload)
}
