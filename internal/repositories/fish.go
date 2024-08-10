package repositories

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/pkg/apihelper"

	"gorm.io/gorm"
)

//go:generate mockgen -source=fish.go -destination=../../test/repositories/mock/fish.go

type FishInterface interface {
	Create(string) (*models.Fish, error)
}

type FishRepository struct {
	db *gorm.DB
}

func NewFishRepository(db *gorm.DB) FishInterface {
	return &FishRepository{
		db: db,
	}
}

func (f FishRepository) Create(name string) (*models.Fish, error) {
	fish := &models.Fish{
		Name: name,
	}

	err := f.db.Create(fish)
	if err.Error != nil {
		return nil, apihelper.GromError(err.Error)
	}

	return fish, nil
}
