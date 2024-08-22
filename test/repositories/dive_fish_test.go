package repositories_test

import (
	"diving-log-book-service/internal/types"
	"diving-log-book-service/test/utils"
	"testing"
)

func TestCreateDiveFish(t *testing.T) {
	utils.DiveToDelete = []uint{}
	utils.DiveFishToDelete = []uint{}

	repo := utils.GetDiveFishRepo()
	diveRepo := utils.GetDiveRepo()

	dive := &types.CreateDivePayload{
		Name:        "test_dive_fish",
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

	utils.DiveToDelete = append(utils.DiveFishToDelete, createdDive.ID)
	utils.DiveFishToDelete = append(utils.DiveFishToDelete, createdDiveFish.ID)
}

func TestReadAllDiveFish(t *testing.T) {
	repo := utils.GetDiveFishRepo()

	_, err := repo.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadOneDiveFish(t *testing.T) {
	repo := utils.GetDiveFishRepo()

	dive, err := repo.ReadOne(utils.DiveFishToDelete[0])
	if err != nil {
		t.Fatal(err)
	}

	if dive.ID != utils.DiveFishToDelete[0] {
		t.Fatal("Should be diveId ", utils.DiveFishToDelete[0], " but got ", dive.ID)
	}
}

func TestDeleteDiveFish(t *testing.T) {
	repo := utils.GetDiveFishRepo()
	diveRepo := utils.GetDiveRepo()

	err := repo.Delete(utils.DiveFishToDelete[0])
	if err != nil {
		t.Fatal(err)
	}

	err = diveRepo.Delete(utils.DiveToDelete[0])
	if err != nil {
		t.Fatal(err)
	}
}
