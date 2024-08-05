package repositories

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/types"
	"fmt"

	"gorm.io/gorm"
)

type DiveInterface interface {
	Create(types.CreateDivePayload) (*models.Dive, error)
}

type DiveRepository struct {
	db *gorm.DB
}

func NewDiveRepository(db *gorm.DB) DiveInterface {
	return &DiveRepository{
		db: db,
	}
}

func (d DiveRepository) Create(payload types.CreateDivePayload) (*models.Dive, error) {
	dive := &models.Dive{
		Name:        payload.Name,
		Depth:       payload.Depth,
		Country:     payload.Country,
		Island:      payload.Island,
		Weight:      payload.Weight,
		Description: payload.Description,
		FishList:    payload.FishList,
		Duration:    payload.Duration,
		UserID:      payload.UserID,
		Media:       payload.Media,
	}

	err := d.db.Create(&dive)
	if err.Error != nil {
		fmt.Println(err.Error)
		return nil, err.Error
	}

	return dive, nil
}
