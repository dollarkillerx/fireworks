package base

import (
	"github.com/dollarkillerx/fireworks/internal/app/backend/storage"
	"github.com/dollarkillerx/fireworks/internal/conf"
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"gorm.io/gorm"

	"log"
	"time"
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
		&models.Configuration{},
	)

	base := &Base{db: postgres}

	// init admin user
	user, err := base.GetUser(conf.GetBackendConfig().BasicAdministrator.Email)
	if err != nil {
		err := base.CreateUser(conf.GetBackendConfig().BasicAdministrator.Email, conf.GetBackendConfig().BasicAdministrator.Name, conf.GetBackendConfig().BasicAdministrator.Password, enum.Admin)
		if err != nil {
			log.Println(err)
		}
	} else {
		err := base.db.Model(&models.User{}).Where("email =?", user.Email).Updates(models.User{
			UserName: conf.GetBackendConfig().BasicAdministrator.Name,
			Password: conf.GetBackendConfig().BasicAdministrator.Password,
		}).Error
		if err != nil {
			log.Println(err)
		}
	}

	go func() {
		for {
			base.updateTaskOldLogs()
			time.Sleep(time.Minute * 10)
		}
	}()

	return base
}

func (b *Base) GetDB() *gorm.DB {
	return b.db
}
