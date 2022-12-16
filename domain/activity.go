package domain

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        int            `gorm:"primarykey auto_increment" json:"id"`
	Email     string         `json:"email"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Todo      *Todo          `gorm:"foreignKey:ActivityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"todo,omitempty"`
}

func (d Activity) TableName() string {
	return "activity"
}


