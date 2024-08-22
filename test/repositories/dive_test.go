package repositories_test

import (
	"diving-log-book-service/internal/types"
	"diving-log-book-service/test/utils"
	"testing"
)

func TestCreateDive(t *testing.T) {
	utils.DiveToDelete = []uint{}
	repo := utils.GetDiveRepo()

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

	utils.DiveToDelete = append(utils.DiveToDelete, createdDive.ID)
}

func TestReadAllDive(t *testing.T) {
	repo := utils.GetDiveRepo()

	_, err := repo.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadOneDive(t *testing.T) {
	repo := utils.GetDiveRepo()

	dive, err := repo.ReadOne(utils.DiveToDelete[0])
	if err != nil {
		t.Fatal(err)
	}

	if dive.ID != utils.DiveToDelete[0] {
		t.Fatal("Should be diveId ", utils.DiveToDelete[0], " but got ", dive.ID)
	}
}

func TestDeleteDive(t *testing.T) {
	repo := utils.GetDiveRepo()

	err := repo.Delete(utils.DiveToDelete[0])
	if err != nil {
		t.Fatal(err)
	}
}
