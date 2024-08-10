package repositories

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/pkg/apihelper"
	"diving-log-book-service/internal/types"

	"gorm.io/gorm"
)

//go:generate mockgen -source=dive-fish.go -destination=../../test/repositories/mock/dive-fish.go

type DiveFishInterface interface {
	Create(*types.CreateFishPayload) (*models.DiveFish, error)
	ReadAll() ([]models.DiveFish, error)
	ReadOne(id uint) (*models.DiveFish, error)
	Delete(id uint) error
}

type DiveFishRepository struct {
	db *gorm.DB
}

func NewDiveFishRepository(db *gorm.DB) DiveFishInterface {
	return &DiveFishRepository{
		db: db,
	}
}

func (fm DiveFishRepository) Create(payload *types.CreateFishPayload) (*models.DiveFish, error) {
	fishMapping := &models.DiveFish{
		DiveID: payload.DiveID,
		FishID: payload.FishID,
	}

	err := fm.db.Create(&fishMapping)
	if err.Error != nil {
		return nil, apihelper.GromError(err.Error)
	}

	return fishMapping, nil
}

func (df DiveFishRepository) ReadAll() ([]models.DiveFish, error) {
	var dives []models.DiveFish

	err := df.db.Find(&dives)
	if err.Error != nil {
		return nil, apihelper.GromError(err.Error)
	}

	return dives, nil
}

func (df DiveFishRepository) ReadOne(id uint) (*models.DiveFish, error) {
	var dive *models.DiveFish

	err := df.db.Where("id = ?", id).Find(&dive)
	if err.Error != nil {
		return nil, apihelper.GromError(err.Error)
	}

	return dive, nil
}

func (df DiveFishRepository) Delete(id uint) error {
	err := df.db.Delete(&models.DiveFish{}, "id = ?", id)
	if err.Error != nil {
		return apihelper.GromError(err.Error)
	}

	return nil
}
