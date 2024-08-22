package repositories

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/pkg/apihelper"
	"diving-log-book-service/internal/types"

	"gorm.io/gorm"
)

//go:generate mockgen -source=media.go -destination=../../test/repositories/mock/media.go

type MediaInterface interface {
	Create(*types.CreateMediaPayload) (*models.Media, error)
	Delete(uint) error
	ReadOne(uint) (*models.Media, error)
}

type MediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) MediaInterface {
	return &MediaRepository{
		db,
	}
}

func (m MediaRepository) Create(payload *types.CreateMediaPayload) (*models.Media, error) {
	media := &models.Media{
		Bucket: payload.Bucket,
		Name:   payload.Key,
		DiveID: payload.DiveID,
	}

	err := m.db.Create(&media)
	if err.Error != nil {
		return nil, apihelper.GromError(err.Error)
	}

	return media, nil
}

func (m MediaRepository) ReadOne(id uint) (*models.Media, error) {
	var media *models.Media

	err := m.db.Where("id = ?", id).Find(&media)
	if err.Error != nil {
		return nil, apihelper.GromError(err.Error)
	}

	return media, nil
}

func (m MediaRepository) Delete(id uint) error {
	err := m.db.Delete(&models.Media{}, "id = ?", id)
	if err.Error != nil {
		return apihelper.GromError(err.Error)
	}

	return nil
}
