package repositories

import (
	"diving-log-book-service/internal/models"
	gormHelper "diving-log-book-service/internal/pkg/gorm"
	"fmt"

	"gorm.io/gorm"
)

//go:generate mockgen -source=country.go -destination=../../test/repositories/mock/country.go

type CountryInterface interface {
	ReadAll(filters []gormHelper.Filter) ([]models.Country, error)
}

type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) CountryInterface {
	return &CountryRepository{
		db,
	}
}

func (c CountryRepository) ReadAll(filters []gormHelper.Filter) ([]models.Country, error) {
	db := c.db
	fmt.Println(filters)
	if filters != nil {
		scopes := gormHelper.CreateFilter(filters)
		db = db.Scopes(scopes)
	}
	var countries []models.Country

	err := db.Find(&countries)
	if err.Error != nil {
		return nil, err.Error
	}

	return countries, nil
}
