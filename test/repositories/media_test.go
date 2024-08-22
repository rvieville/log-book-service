package repositories_test

import (
	"diving-log-book-service/internal/types"
	"diving-log-book-service/test/utils"
	"testing"
)

func TestCreateMedia(t *testing.T) {
	utils.DiveToDelete = []uint{}
	repo := utils.GetMediaRepo()
	diveRepo := utils.GetDiveRepo()

	dive := &types.CreateDivePayload{
		Name:        "test_media",
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

	utils.DiveToDelete = append(utils.DiveToDelete, createdDive.ID)

	media := &types.CreateMediaPayload{
		Bucket: "dive",
		Key:    "test.png",
		DiveID: createdDive.ID,
	}

	createMedia, err := repo.Create(media)
	if err != nil {
		t.Fatal(err)
	}

	utils.MediaToDelete = append(utils.MediaToDelete, createMedia.ID)
}

func TestReadOneMedia(t *testing.T) {
	repo := utils.GetMediaRepo()

	_, err := repo.ReadOne(utils.MediaToDelete[0])
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteMedia(t *testing.T) {
	repo := utils.GetMediaRepo()
	diveRepo := utils.GetDiveRepo()

	err := repo.Delete(utils.MediaToDelete[0])
	if err != nil {
		t.Fatal(err)
	}

	err = diveRepo.Delete(utils.DiveToDelete[0])
	if err != nil {
		t.Fatal(err)
	}
}
