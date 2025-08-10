package repositories

import (
	"attendance-app-api/models"
	"errors"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAll() (*[]models.User, error)
	FindById(userId uint) (*models.User, error)
	Create(newUser models.User) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() (*[]models.User, error) {
	users := []models.User{}
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (r *UserRepository) FindById(userId uint) (*models.User, error) {
	user := models.User{}
	result := r.db.First(&user, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Create(newUser models.User) (*models.User, error) {
	result := r.db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil
}
