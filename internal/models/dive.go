package models

import "time"

type Dive struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Depth       float32   `json:"depth"`
	Country     string    `json:"country"`
	Island      string    `json:"island"`
	Weight      float32   `json:"weight"`
	Description string    `json:"description"`
	Fishes      []Fish    `gorm:"many2many:dive_fish"`
	Medias      []Media   `json:"medias"`
	Duration    float32   `json:"duration"`
	UserID      *uint     `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (Dive) TableName() string {
	return "dive"
}
