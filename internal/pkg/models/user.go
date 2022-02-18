package models

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string    `gorm:"uniqueKey;type:varchar(32)"`
	UserName string    `gorm:"type:varchar(250)"`
	Password string    `gorm:"type:varchar(250)"`
	Role     enum.Role `gorm:"type:varchar(250)"`
}
