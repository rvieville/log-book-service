package models

import (
	"gorm.io/gorm"
)

type Dive struct {
	gorm.Model

	Name        string  `json:"name"`
	Depth       float32 `json:"depth"`
	Country     string  `json:"country"`
	Island      string  `json:"island"`
	Weight      float32 `json:"weight"`
	Description string  `json:"description"`
	Fishes      []Fish  `gorm:"many2many:dive_fish"`
	Duration    float32 `json:"duration"`
	UserID      *uint   `json:"userId"`
}

func (Dive) TableName() string {
	return "dive"
}
