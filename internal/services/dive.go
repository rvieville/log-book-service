package services

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
)

type DiveService struct {
	diveRepo        repositories.DiveInterface
	diveFishService *DiveFishService
}

type DiveServiceConfig struct {
	DiveRepo        repositories.DiveInterface
	DiveFishService *DiveFishService
}

func NewDiveService(config *DiveServiceConfig) *DiveService {
	return &DiveService{
		diveRepo:        config.DiveRepo,
		diveFishService: config.DiveFishService,
	}
}

func (fm DiveService) Create(payload *types.CreateDivePayload) (*models.Dive, error) {
	dive, err := fm.diveRepo.Create(payload)
	if err != nil {
		return nil, err
	}

	for _, fish := range payload.Fishes {
		fm.diveFishService.Create(&types.CreateFishPayload{
			DiveID: dive.ID,
			FishID: fish,
		})
	}
	// TODO: create media

	return dive, nil
}

func (fm DiveService) ReadAll() ([]models.Dive, error) {
	dives, err := fm.diveRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return dives, nil
}
