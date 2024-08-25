package repositories

import (
	"diving-log-book-service/internal/models"
	gormHelper "diving-log-book-service/internal/pkg/gorm"

	"gorm.io/gorm"
)

//go:generate mockgen -source=island.go -destination=../../test/repositories/mock/island.go

type IslandInterface interface {
	ReadAll(filters []gormHelper.Filter) ([]models.Island, error)
}

type IslandRepository struct {
	db *gorm.DB
}

func NewIslandRepository(db *gorm.DB) IslandInterface {
	return &IslandRepository{
		db,
	}
}

func (i IslandRepository) ReadAll(filters []gormHelper.Filter) ([]models.Island, error) {
	db := i.db
	if filters != nil {
		scope := gormHelper.CreateFilter(filters)
		db = db.Scopes(scope)
	}
	var islands []models.Island

	err := db.Find(&islands)
	if err.Error != nil {
		return nil, err.Error
	}

	return islands, nil
}
