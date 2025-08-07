package models

type Shift struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	StartTime   string `gorm:"not null"`
	EndTime     string `gorm:"not null"`
	WorkContent string `gorm:"not null"`
	Issues      string `gorm:"not null"`
}
