package models

import "time"

type Media struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	Bucket    string
	Name      string
	Url       string `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Media) TableName() string {
	return "media"
}
