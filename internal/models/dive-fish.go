package models

import "time"

type DiveFish struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	DiveID    uint
	FishID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (DiveFish) TableName() string {
	return "dive_fish"
}
