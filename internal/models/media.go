package models

import (
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model

	DiveID int32
}

func (Media) TableName() string {
	return "media"
}
