package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://admin:admin@0.0.0.0:5432/log-book"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
