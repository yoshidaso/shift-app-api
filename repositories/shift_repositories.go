package repositories

import (
	"attendance-app-api/models"
	"errors"
)

type IShiftRepository interface {
	FindAll() (*[]models.Shift, error)
	FindById(shiftId uint) (*models.Shift, error)
}

type ShiftMemoryRepository struct {
	shifts []models.Shift
}

func NewShiftMemoryRepository(shifts []models.Shift) IShiftRepository {
	return &ShiftMemoryRepository{shifts: shifts}
}

func (r *ShiftMemoryRepository) FindAll() (*[]models.Shift, error) {
	return &r.shifts, nil
}

func (r *ShiftMemoryRepository) FindById(shiftId uint) (*models.Shift, error) {
	for _, v := range r.shifts {
		if v.ID == shiftId {
			return &v, nil
		}
	}
	return nil, errors.New("shift not found")
}
