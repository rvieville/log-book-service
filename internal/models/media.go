package models

import (
	"time"
)

type Media struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Bucket    string    `json:"bucket"`
	Name      string    `json:"name"`
	DiveID    uint      `json:"dive_id"`
	Url       string    `gorm:"-" json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_At"`
}

func (m Media) TableName() string {
	return "media"
}
