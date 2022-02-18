package request

import "github.com/dollarkillerx/fireworks/internal/pkg/enum"

type WebLogin struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type AuthModel struct {
	Email string    `json:"email" validate:"required"`
	Name  string    `json:"name" validate:"required"`
	Role  enum.Role `json:"role"`
}
