package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u User) TableName() string {
	return "user"
}
