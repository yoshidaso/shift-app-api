package services

import (
	"attendance-app-api/dto"
	"attendance-app-api/models"
	"attendance-app-api/repositories"
)

type IUserService interface {
	FindAll() (*[]models.User, error)
	FindById(userId uint) (*models.User, error)
	Create(createUserRequest dto.CreateUserRequest) (*models.User, error)
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) FindAll() (*[]models.User, error) {
	return s.userRepository.FindAll()
}

func (s *UserService) FindById(userId uint) (*models.User, error) {
	return s.userRepository.FindById(userId)
}

func (s *UserService) Create(createUserRequest dto.CreateUserRequest) (*models.User, error) {
	newUser := models.User{
		Name:  createUserRequest.Name,
		Email: createUserRequest.Email,
	}
	return s.userRepository.Create(newUser)
}
