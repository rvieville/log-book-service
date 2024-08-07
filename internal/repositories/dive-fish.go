package repositories

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/types"

	"gorm.io/gorm"
)

//go:generate mockgen -source=dive-fish.go -destination=../../test/repositories/mock/dive-fish.go

type DiveFishInterface interface {
	Create(types.CreateFishMappingPayload) (*models.DiveFish, error)
}

type DiveFishRepository struct {
	db *gorm.DB
}

func NewDiveFishRepository(db *gorm.DB) DiveFishInterface {
	return &DiveFishRepository{
		db: db,
	}
}

func (fm DiveFishRepository) Create(payload types.CreateFishMappingPayload) (*models.DiveFish, error) {
	fishMapping := &models.DiveFish{
		DiveID: payload.DiveID,
		FishID: payload.FishID,
	}

	err := fm.db.Create(&fishMapping)
	if err.Error != nil {
		return nil, err.Error
	}

	return fishMapping, nil
}
