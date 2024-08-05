package repositories

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/types"

	"gorm.io/gorm"
)

type FishMappingInterface interface {
	Create(types.CreateFishMappingPayload) (*models.FishMapping, error)
}

type FishMappingRepository struct {
	db *gorm.DB
}

func NewFishMappingRepository(db *gorm.DB) FishMappingInterface {
	return &FishMappingRepository{
		db: db,
	}
}

func (fm FishMappingRepository) Create(payload types.CreateFishMappingPayload) (*models.FishMapping, error) {
	fishMapping := &models.FishMapping{
		DiveID: payload.DiveID,
		FishID: payload.FishID,
	}

	err := fm.db.Create(&fishMapping)
	if err.Error != nil {
		return nil, err.Error
	}

	return fishMapping, nil
}
