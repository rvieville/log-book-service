package models

import "time"

type Fish struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Fish) TableName() string {
	return "fish"
}
