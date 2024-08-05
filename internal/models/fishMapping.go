package models

import (
	"gorm.io/gorm"
)

type FishMapping struct {
	gorm.Model

	DiveID int32
	FishID int32
	Dive   *Dive `gorm:"foreignKey:DiveID"`
	Fish   *Fish `gorm:"foreignKey:FishID"`
}

func (FishMapping) TableName() string {
	return "fish_mapping"
}
