package models

import "time"

type User struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	DisplayName string
	Email       string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u User) TableName() string {
	return "user"
}
