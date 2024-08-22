package repositories_test

import (
	"diving-log-book-service/internal/repositories"
	"diving-log-book-service/internal/types"
	"diving-log-book-service/test/utils"
	"testing"
)

var userToDelete []uint

func getUserRepo() repositories.UserInterface {
	return repositories.NewUserRepository(utils.ConnectDB())
}

func TestCreateUser(t *testing.T) {
	userToDelete = []uint{}
	repo := getUserRepo()

	user := &types.CreateUserPayload{
		DisplayName: "rafael vieville",
		Email:       "rafael.vieville@gmail.com",
		Password:    "password",
	}

	createdUser, err := repo.Create(user)
	if err != nil {
		t.Fatal(err)
	}

	userToDelete = append(userToDelete, createdUser.ID)
}

func TestReadAllUser(t *testing.T) {
	repo := getUserRepo()

	_, err := repo.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadOneUser(t *testing.T) {
	repo := getUserRepo()

	dive, err := repo.ReadOne(userToDelete[0])
	if err != nil {
		t.Fatal(err)
	}

	if dive.ID != userToDelete[0] {
		t.Fatal("Should be diveId ", userToDelete[0], " but got ", dive.ID)
	}
}

func TestDeleteUser(t *testing.T) {
	repo := getUserRepo()

	err := repo.Delete(userToDelete[0])
	if err != nil {
		t.Fatal(err)
	}
}
