package dive

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
)

type Dive struct {
	repository repositories.DiveInterface
}

func New(repository repositories.DiveInterface) *Dive {
	return &Dive{
		repository,
	}
}

func (fm Dive) Create(payload types.CreateDivePayload) (*models.Dive, error) {
	return fm.repository.Create(payload)
}
