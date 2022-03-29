package base

import (
	"github.com/dollarkillerx/fireworks/internal/app/backend/storage"
	"github.com/dollarkillerx/fireworks/internal/conf"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"gorm.io/gorm"
)

type Base struct {
	db *gorm.DB
}

func NewBase() storage.Interface {
	postgres, err := utils.InitPostgres(conf.GetBackendConfig().PostgresConfig, conf.GetBackendConfig().Debug)
	if err != nil {
		return nil
	}

	postgres.AutoMigrate(
		&models.User{},
		&models.Agent{},
		&models.TaskLog{},
		&models.Task{},
		&models.Subtasks{},
	)

	return &Base{db: postgres}
}

func (b *Base) GetDB() *gorm.DB {
	return b.db
}
