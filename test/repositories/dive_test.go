package repositories_test

import (
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

var diveToDelete []uint
var diveFishToDelete []uint

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db.Connection()

	m.Run()
}

func TestCreateDive(t *testing.T) {
	diveToDelete = []uint{}
	diveFishToDelete = []uint{}
	repo := repositories.NewDiveRepository(db.DB)

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

	diveToDelete = append(diveToDelete, createdDive.ID)
}

func TestReadAllDive(t *testing.T) {
	repo := repositories.NewDiveRepository(db.DB)

	_, err := repo.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadOneDive(t *testing.T) {
	repo := repositories.NewDiveRepository(db.DB)

	dive, err := repo.ReadOne(diveToDelete[0])
	if err != nil {
		t.Fatal(err)
	}

	if dive.ID != diveToDelete[0] {
		t.Fatal("Should be diveId ", diveToDelete[0], " but got ", dive.ID)
	}
}

func TestDeleteDive(t *testing.T) {
	repo := repositories.NewDiveRepository(db.DB)

	err := repo.Delete(diveToDelete[0])
	if err != nil {
		t.Fatal(err)
	}
}
