package models

import "time"

type DiveMedia struct {
	ID        uint `gorm:"primaryKey"`
	DiveID    uint
	MediaID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (dm DiveMedia) TableName() string {
	return "dive_media"
}
