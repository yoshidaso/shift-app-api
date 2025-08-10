package services

import (
	"attendance-app-api/dto"
	"attendance-app-api/models"
	"attendance-app-api/repositories"
	"time"
)

type IShiftService interface {
	FindAll() (*[]models.Shifts, error)
	FindById(shiftId uint) (*models.Shifts, error)
	Create(createShiftInput dto.CreateShiftRequest) (*models.Shifts, error)
}

type ShiftService struct {
	repository repositories.IShiftRepository
}

func NewShiftService(repository repositories.IShiftRepository) IShiftService {
	return &ShiftService{repository: repository}
}

func (s *ShiftService) FindAll() (*[]models.Shifts, error) {
	return s.repository.FindAll()
}

func (s *ShiftService) FindById(shiftId uint) (*models.Shifts, error) {
	return s.repository.FindById(shiftId)
}

func (s *ShiftService) Create(createShitRequest dto.CreateShiftRequest) (*models.Shifts, error) {
	today := time.Now().Format("2006-01-02")
	startTime, err := time.Parse("2006-01-02 15:04", today+" "+createShitRequest.StartTime)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("2006-01-02 15:04", today+" "+createShitRequest.EndTime)
	if err != nil {
		return nil, err
	}

	newShift := models.Shifts{
		UserID:      createShitRequest.UserID,
		StartAt:     startTime,
		EndAt:       endTime,
		WorkContent: createShitRequest.WorkContent,
		Issues:      createShitRequest.Issues,
	}
	return s.repository.Create(newShift)
}
