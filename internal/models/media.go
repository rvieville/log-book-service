package models

import (
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model

	DiveID uint
	Dive   *Dive `gorm:"foreignKey:DiveID"`
}

func (Media) TableName() string {
	return "media"
}
