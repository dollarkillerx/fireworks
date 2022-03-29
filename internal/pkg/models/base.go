package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string         `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
