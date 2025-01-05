package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
