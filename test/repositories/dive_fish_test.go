package repositories_test

import (
	"diving-log-book-service/internal/db"
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
	"testing"
)

func TestCreateDiveFish(t *testing.T) {
	diveToDelete = []uint{}
	diveFishToDelete = []uint{}

	repo := repositories.NewDiveFishRepository(db.DB)
	diveRepo := repositories.NewDiveRepository(db.DB)

	dive := &types.CreateDivePayload{
		Name:        "bonjour",
		Depth:       14,
		Country:     "France",
		Island:      "Reunion",
		Weight:      7,
		Description: "couocu",
		Duration:    41.27,
	}

	createdDive, err := diveRepo.Create(dive)
	if err != nil {
		t.Fatal(err)
	}

	diveFish := &types.CreateFishPayload{
		DiveID: createdDive.ID,
		FishID: 1,
	}

	createdDiveFish, err := repo.Create(diveFish)
	if err != nil {
		t.Fatal(err)
	}

	diveToDelete = append(diveFishToDelete, createdDive.ID)
	diveFishToDelete = append(diveFishToDelete, createdDiveFish.ID)
}

func TestReadAllDiveFish(t *testing.T) {
	repo := repositories.NewDiveFishRepository(db.DB)

	_, err := repo.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadOneDiveFish(t *testing.T) {
	repo := repositories.NewDiveFishRepository(db.DB)

	dive, err := repo.ReadOne(diveFishToDelete[0])
	if err != nil {
		t.Fatal(err)
	}

	if dive.ID != diveFishToDelete[0] {
		t.Fatal("Should be diveId ", diveFishToDelete[0], " but got ", dive.ID)
	}
}

func TestDeleteDiveFish(t *testing.T) {
	repo := repositories.NewDiveFishRepository(db.DB)
	diveRepo := repositories.NewDiveRepository(db.DB)

	err := repo.Delete(diveFishToDelete[0])
	if err != nil {
		t.Fatal(err)
	}

	err = diveRepo.Delete(diveToDelete[0])
	if err != nil {
		t.Fatal(err)
	}
}
