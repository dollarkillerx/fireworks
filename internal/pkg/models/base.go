package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string         `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
