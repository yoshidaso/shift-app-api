package repositories

import (
	"attendance-app-api/models"
	"errors"

	"gorm.io/gorm"
)

type IShiftRepository interface {
	FindAll() (*[]models.Shift, error)
	FindById(shiftId uint) (*models.Shift, error)
	Create(newShift models.Shift) (*models.Shift, error)
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

func (r *ShiftMemoryRepository) Create(newShift models.Shift) (*models.Shift, error) {
	newShift.ID = uint(len(r.shifts) + 1)
	r.shifts = append(r.shifts, newShift)
	return &newShift, nil
}

type ShiftRepository struct {
	db *gorm.DB
}

func NewShiftRepository(db *gorm.DB) IShiftRepository {
	return &ShiftRepository{db: db}
}

func (r *ShiftRepository) FindAll() (*[]models.Shift, error) {
	shifts := []models.Shift{}
	result := r.db.Find(&shifts)
	if result.Error != nil {
		return nil, result.Error
	}
	return &shifts, nil
}

func (r *ShiftRepository) FindById(shiftId uint) (*models.Shift, error) {
	shift := models.Shift{}
	result := r.db.First(&shift, shiftId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("shift not found")
		}
		return nil, result.Error
	}
	return &shift, nil
}

func (r *ShiftRepository) Create(newShift models.Shift) (*models.Shift, error) {
	result := r.db.Create(&newShift)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newShift, nil
}
