package repositories

import (
	"diving-log-book-service/internal/models"
	"diving-log-book-service/internal/pkg/crypto"
	"diving-log-book-service/internal/types"

	"gorm.io/gorm"
)

//go:generate mockgen -source=user.go -destination=../../test/repositories/mock/user.go

type UserInterface interface {
	Create(*types.CreateUserPayload) (*models.User, error)
	ReadAll() ([]models.User, error)
	ReadOne(id uint) (*models.User, error)
	Delete(id uint) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserInterface {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) Create(payload *types.CreateUserPayload) (*models.User, error) {
	hashedPassword, passErr := crypto.HashPassword(payload.Password)
	if passErr != nil {
		return nil, passErr
	}

	user := &models.User{
		DisplayName: payload.DisplayName,
		Email:       payload.Email,
		Password:    hashedPassword,
	}

	err := u.db.Create(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return user, nil
}

func (u UserRepository) ReadAll() ([]models.User, error) {
	var users []models.User

	err := u.db.Find(&users)
	if err.Error != nil {
		return nil, err.Error
	}

	return users, nil
}

func (u UserRepository) ReadOne(id uint) (*models.User, error) {
	var user *models.User

	err := u.db.Where("id = ?", id).Find(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return user, nil
}

func (u UserRepository) Delete(id uint) error {
	err := u.db.Delete(&models.User{}, "id = ?", id)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
