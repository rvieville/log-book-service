package models

import "time"

type DiveFish struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	DiveID    uint      `json:"dive_id"`
	FishID    uint      `json:"fish_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DiveFish) TableName() string {
	return "dive_fish"
}
