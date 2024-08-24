package main

import (
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

type Test struct {
	Country string
	Islands []string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connection()
	fmt.Println("connect success")

	err = db.DB.AutoMigrate(&models.Dive{}, &models.DiveFish{}, &models.Fish{}, &models.Media{}, &models.User{}, &models.Country{}, &models.Island{})
	if err != nil {
		panic(err.Error())
	}

	// file, err := os.ReadFile("cmd/migration/place.json")
	// if err != nil {
	// 	panic(err)
	// }
	// var places []Test

	// json.Unmarshal(file, &places)

	// for _, location := range places {
	// 	country := &models.Country{
	// 		Name: location.Country,
	// 		Slug: strings.ToLower(strings.ReplaceAll(location.Country, " ", "-")),
	// 	}
	// 	db.DB.Create(country)

	// 	for _, island := range location.Islands {
	// 		island := &models.Island{
	// 			Name:      island,
	// 			CountryID: country.ID,
	// 			Slug:      strings.ToLower(strings.ReplaceAll(island, " ", "-")),
	// 		}
	// 		db.DB.Create(island)
	// 	}
	// }
	// fmt.Println("Tables reated")
	// file, err := os.Open("cmd/migration/fish.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// scanner := bufio.NewScanner(file)
	// fishReo := repositories.NewFishRepository(db.DB)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// 	fishReo.Create(scanner.Text())
	// }
}
