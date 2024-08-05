package models

import (
	"gorm.io/gorm"
)

type Dive struct {
	gorm.Model

	Name        string
	Depth       float32
	Country     string
	Island      string
	Weight      float32
	Description string
	FishList    *[]FishMapping
	Duration    float32
	UserID      *int32
	Media       *[]Media
}

func (Dive) TableName() string {
	return "dive"
}
