package utils

import (
	"diving-log-book-service/internal/repositories"

	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DiveToDelete     []uint
	MediaToDelete    []uint
	DiveFishToDelete []uint
	UploadID         *string
	Parts            []*s3.CompletedPart
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://admin:admin@0.0.0.0:5432/log-book"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func GetMediaRepo() repositories.MediaInterface {
	return repositories.NewMediaRepository(ConnectDB())
}

func GetDiveRepo() repositories.DiveInterface {
	return repositories.NewDiveRepository(ConnectDB())
}

func GetDiveFishRepo() repositories.DiveFishInterface {
	return repositories.NewDiveFishRepository(ConnectDB())
}

func GetCountryRepo() repositories.CountryInterface {
	return repositories.NewCountryRepository(ConnectDB())
}
