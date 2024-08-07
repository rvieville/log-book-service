package repositories

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/types"
	"fmt"

	"gorm.io/gorm"
)

//go:generate mockgen -source=dive.go -destination=../../test/repositories/mock/dive.go

type DiveInterface interface {
	Create(*types.CreateDivePayload) (*models.Dive, error)
	ReadAll() ([]models.Dive, error)
	ReadOne(id uint) (*models.Dive, error)
	Delete(id uint) error
}

type DiveRepository struct {
	db *gorm.DB
}

func NewDiveRepository(db *gorm.DB) DiveInterface {
	return &DiveRepository{
		db: db,
	}
}

func (d DiveRepository) Create(payload *types.CreateDivePayload) (*models.Dive, error) {
	dive := &models.Dive{
		Name:        payload.Name,
		Depth:       payload.Depth,
		Country:     payload.Country,
		Island:      payload.Island,
		Weight:      payload.Weight,
		Description: payload.Description,
		Duration:    payload.Duration,
		UserID:      payload.UserID,
	}

	err := d.db.Create(&dive)
	if err.Error != nil {
		fmt.Println(err.Error)
		return nil, err.Error
	}

	return dive, nil
}

func (d DiveRepository) ReadAll() ([]models.Dive, error) {
	var dives []models.Dive

	err := d.db.Preload("Fishes").Find(&dives)
	if err.Error != nil {
		return nil, err.Error
	}

	return dives, nil
}

func (d DiveRepository) ReadOne(id uint) (*models.Dive, error) {
	var dive *models.Dive

	err := d.db.Where("id = ?", id).Find(&dive)
	if err.Error != nil {
		return nil, err.Error
	}

	return dive, nil
}

func (d DiveRepository) Delete(id uint) error {
	err := d.db.Delete(&models.Dive{}, "id = ?", id)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
