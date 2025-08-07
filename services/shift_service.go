package services

import (
	"attendance-app-api/dto"
	"attendance-app-api/models"
	"attendance-app-api/repositories"
)

type IShiftService interface {
	FindAll() (*[]models.Shift, error)
	FindById(shiftId uint) (*models.Shift, error)
	Create(createShiftInput dto.CreateShiftRequest) (*models.Shift, error)
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

func (s *ShiftService) Create(createShitRequest dto.CreateShiftRequest) (*models.Shift, error) {
	newShift := models.Shift{
		UserID:      createShitRequest.UserID,
		StartTime:   createShitRequest.StartTime,
		EndTime:     createShitRequest.EndTime,
		WorkContent: createShitRequest.WorkContent,
		Issues:      createShitRequest.Issues,
	}
	return s.repository.Create(newShift)
}
