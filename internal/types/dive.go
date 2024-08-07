package types

import (
	"diving-log-book-service/internal/models"
)

type CreateDivePayload struct {
	Name        string  `validate:"required"`
	Depth       float32 `validate:"required"`
	Country     string  `validate:"required"`
	Island      string  `validate:"required"`
	Weight      float32 `validate:"required"`
	Description string  `validate:"required"`
	Fishes      []uint  `validate:"required"`
	Duration    float32 `validate:"required"`
	UserID      *uint
	Media       *[]models.Media
}
