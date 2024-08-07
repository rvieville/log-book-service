package repositories_test

import (
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

var elemToDelete []uint

func GetRepo() repositories.DiveInterface {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db.Connection()
	return repositories.NewDiveRepository(db.DB)
}

func TestCreateDive(t *testing.T) {
	repo := GetRepo()

	dive := &types.CreateDivePayload{
		Name:        "bonjour",
		Depth:       14,
		Country:     "France",
		Island:      "Reunion",
		Weight:      7,
		Description: "couocu",
		Duration:    41.27,
	}

	createdDive, err := repo.Create(dive)
	if err != nil {
		t.Fatal(err)
	}

	elemToDelete = append(elemToDelete, createdDive.ID)
}

func TestReadAll(t *testing.T) {
	repo := GetRepo()

	_, err := repo.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadOne(t *testing.T) {
	repo := GetRepo()

	dive, err := repo.ReadOne(elemToDelete[0])
	if err != nil {
		t.Fatal(err)
	}

	if dive.ID != elemToDelete[0] {
		t.Fatal("Should be diveId ", elemToDelete[0], " but got ", dive.ID)
	}
}

func TestDelete(t *testing.T) {
	repo := GetRepo()

	fmt.Println(elemToDelete)

	err := repo.Delete(elemToDelete[0])
	if err != nil {
		t.Fatal(err)
	}
}
