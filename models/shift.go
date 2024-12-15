package models

type Shift struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	StartTime string `gorm:"not null"`
	EndTime   string `gorm:"not null"`
}
