package models

import "time"

type Users struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP" json:"createdAt"`
}

func (Users) TableName() string {
	return "Users"
}
