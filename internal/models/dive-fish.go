package models

import (
	"gorm.io/gorm"
)

type DiveFish struct {
	gorm.Model

	DiveID uint
	FishID uint
}

func (DiveFish) TableName() string {
	return "dive_fish"
}
