package repositories

import (
	"attendance-app-api/models"
	"errors"

	"gorm.io/gorm"
)

type IShiftRepository interface {
	FindAll() (*[]models.Shifts, error)
	FindById(shiftId uint) (*models.Shifts, error)
	Create(newShift models.Shifts) (*models.Shifts, error)
	Delete(shiftId uint) error
}

type ShiftMemoryRepository struct {
	shifts []models.Shifts
}

type ShiftRepository struct {
	db *gorm.DB
}

func NewShiftRepository(db *gorm.DB) IShiftRepository {
	return &ShiftRepository{db: db}
}

func (r *ShiftRepository) FindAll() (*[]models.Shifts, error) {
	shifts := []models.Shifts{}
	result := r.db.Find(&shifts)
	if result.Error != nil {
		return nil, result.Error
	}
	return &shifts, nil
}

func (r *ShiftRepository) FindById(shiftId uint) (*models.Shifts, error) {
	shift := models.Shifts{}
	result := r.db.First(&shift, shiftId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("shift not found")
		}
		return nil, result.Error
	}
	return &shift, nil
}

func (r *ShiftRepository) Create(newShift models.Shifts) (*models.Shifts, error) {
	result := r.db.Create(&newShift)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newShift, nil
}

func (r *ShiftRepository) Delete(shiftId uint) error {
	shift := models.Shifts{}
	result := r.db.First(&shift, shiftId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("shift not found")
		}
		return result.Error
	}

	result = r.db.Delete(&shift)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
