package models

import "time"

type Media struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	DiveID    uint
	Dive      *Dive `gorm:"foreignKey:DiveID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Media) TableName() string {
	return "media"
}
