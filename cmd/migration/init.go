package main

import (
	"bufio"
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/repositories"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connection()
	fmt.Println("connect success")

	err = db.DB.AutoMigrate(&models.Dive{}, &models.FishMapping{}, &models.Fish{}, &models.Media{})
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println("Tables reated")
	file, err := os.Open("cmd/migration/fish.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	fishReo := repositories.NewFishRepository(db.DB)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fishReo.Create(scanner.Text())
	}
}
