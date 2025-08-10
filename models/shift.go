package models

import "time"

type Shifts struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"column:user_id;not null" json:"userId"`
	StartAt     time.Time `gorm:"column:startAt;not null" json:"startAt"`
	EndAt       time.Time `gorm:"column:endAt;not null" json:"endAt"`
	WorkContent string    `gorm:"column:workContent;not null" json:"workContent"`
	Issues      string    `gorm:"column:issues;not null" json:"issues"`
}

func (Shifts) TableName() string {
	return "Shifts"
}
