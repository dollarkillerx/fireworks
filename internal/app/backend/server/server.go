package server

import (
	"github.com/dollarkillerx/fireworks/internal/conf"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Backend struct {
	conf *conf.BackendConfig
	app  *gin.Engine
	db   *gorm.DB
}

func NewBackend(conf *conf.BackendConfig, db *gorm.DB) *Backend {
	db.AutoMigrate(&models.User{})
	return &Backend{
		conf: conf,
		db:   db,
		app:  gin.New(),
	}
}

func (b *Backend) Run() error {
	b.app.Use(gin.Recovery())
	if b.conf.Debug {
		b.app.Use(gin.Logger())
	}
	b.router()

	return nil
}
