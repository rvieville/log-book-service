package types

import (
	"diving-log-book-service/internal/models"
	"errors"
)

type CreateDivePayload struct {
	Name        string
	Depth       float32
	Country     string
	Island      string
	Weight      float32
	Description string
	FishList    *[]models.FishMapping
	Duration    float32
	UserID      *int32
	Media       *[]models.Media
}

func (d CreateDivePayload) Validate() error {
	if d.Name == "" {
		return errors.New("name is required")
	} else if d.Depth == 0 {
		return errors.New("depth is required")
	} else if d.Country == "" {
		return errors.New("country is required")
	} else if d.Island == "" {
		return errors.New("island is required")
	} else if d.Description == "" {
		return errors.New("description is required")
	} else if d.Duration == 0 {
		return errors.New("duration is required")
	}

	return nil
}
