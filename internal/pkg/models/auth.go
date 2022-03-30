package models

import "github.com/dollarkillerx/fireworks/internal/pkg/enum"

type AuthToken struct {
	Email string    `json:"email"`
	Name  string    `json:"name"`
	Role  enum.Role `json:"role"`
}
