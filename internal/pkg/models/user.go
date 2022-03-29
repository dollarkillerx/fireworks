package models

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"

	"time"
)

type User struct {
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`

	Email    string    `gorm:"uniqueKey;type:varchar(32)" json:"email"`
	UserName string    `gorm:"type:varchar(250)" json:"user_name"`
	Password string    `gorm:"type:varchar(250)" json:"password"`
	Role     enum.Role `gorm:"type:varchar(250)" json:"role"`
}
