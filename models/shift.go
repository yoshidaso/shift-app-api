package models

import "time"

type Shifts struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName    string    `gorm:"column:user_name;not null" json:"userName"`
	StartAt     time.Time `gorm:"column:start_at;not null" json:"startAt"`
	EndAt       time.Time `gorm:"column:end_at;not null" json:"endAt"`
	WorkContent string    `gorm:"column:work_content;not null" json:"workContent"`
	Issues      string    `gorm:"column:issues;not null" json:"issues"`
}

func (Shifts) TableName() string {
	return "Shifts"
}
