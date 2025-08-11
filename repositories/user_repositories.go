package repositories

import (
	"attendance-app-api/models"
	"errors"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAll() (*[]models.Users, error)
	FindById(userId uint) (*models.Users, error)
	Create(newUser models.Users) (*models.Users, error)
	Delete(userId uint) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() (*[]models.Users, error) {
	users := []models.Users{}
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (r *UserRepository) FindById(userId uint) (*models.Users, error) {
	user := models.Users{}
	result := r.db.First(&user, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Create(newUser models.Users) (*models.Users, error) {
	result := r.db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil
}

func (r *UserRepository) Delete(userId uint) error {
	user := models.Users{}
	result := r.db.First(&user, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return result.Error
	}

	result = r.db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
