package storage

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"gorm.io/gorm"
)

type Interface interface {
	GetDB() *gorm.DB

	user
	agent
}

type user interface {
	CreateUser(email string, name string, password string, role enum.Role) error
	GetUser(email string) (*models.User, error)
}

type agent interface {
	AddAgent(agentName string, agentUrl string, workspace string, description string) error
}
