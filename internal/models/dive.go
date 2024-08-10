package models

import "time"

type Dive struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Depth       float32 `json:"depth"`
	Country     string  `json:"country"`
	Island      string  `json:"island"`
	Weight      float32 `json:"weight"`
	Description string  `json:"description"`
	Fishes      []Fish  `gorm:"many2many:dive_fish"`
	Medias      []Media `gorm:"many2many:dive_media"`
	Duration    float32 `json:"duration"`
	UserID      *uint   `json:"userId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Dive) TableName() string {
	return "dive"
}
