package models

import "gorm.io/gorm"

type Fish struct {
	gorm.Model

	Name string
}

func (Fish) TableName() string {
	return "fish"
}
