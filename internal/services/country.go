package services

import (
	"diving-log-book-service/internal/models"
	gormHelper "diving-log-book-service/internal/pkg/gorm"
	"diving-log-book-service/internal/repositories"
)

type CountryService struct {
	countryRepo repositories.CountryInterface
}

func NewCountryService(countryRepo repositories.CountryInterface) *CountryService {
	return &CountryService{
		countryRepo,
	}
}

func (c CountryService) ReadAll(filters []gormHelper.Filter) ([]models.Country, error) {
	countries, err := c.countryRepo.ReadAll(filters)
	if err != nil {
		return nil, err
	}

	return countries, nil
}
