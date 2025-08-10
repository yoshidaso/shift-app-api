package services

import (
	"attendance-app-api/dto"
	"attendance-app-api/models"
	"attendance-app-api/repositories"
)

type IUserService interface {
	FindAll() (*[]models.Users, error)
	FindById(userId uint) (*models.Users, error)
	Create(createUserRequest dto.CreateUserRequest) (*models.Users, error)
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) FindAll() (*[]models.Users, error) {
	return s.userRepository.FindAll()
}

func (s *UserService) FindById(userId uint) (*models.Users, error) {
	return s.userRepository.FindById(userId)
}

func (s *UserService) Create(createUserRequest dto.CreateUserRequest) (*models.Users, error) {
	newUser := models.Users{
		Name:  createUserRequest.Name,
		Email: createUserRequest.Email,
	}
	return s.userRepository.Create(newUser)
}
