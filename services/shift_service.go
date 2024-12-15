package services

import (
	"attendance-app-api/models"
	"attendance-app-api/repositories"
)

type IShiftService interface {
	FindAll() (*[]models.Shift, error)
	FindById(shiftId uint) (*models.Shift, error)
}

type ShiftService struct {
	repository repositories.IShiftRepository
}

func NewShiftService(repository repositories.IShiftRepository) IShiftService {
	return &ShiftService{repository: repository}
}

func (s *ShiftService) FindAll() (*[]models.Shift, error) {
	return s.repository.FindAll()
}

func (s *ShiftService) FindById(shiftId uint) (*models.Shift, error) {
	return s.repository.FindById(shiftId)
}
