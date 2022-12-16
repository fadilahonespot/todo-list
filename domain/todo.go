package domain

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID         int            `gorm:"primarykey auto_increment" json:"id"`
	ActivityID int            `json:"activity_group_id"`
	Title      string         `json:"title"`
	IsActive   string         `json:"is_active"`
	Priority   string         `json:"priority"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (d Todo) TableName() string {
	return "todo"
}
