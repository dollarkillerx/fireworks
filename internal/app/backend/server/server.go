package server

import (
	"github.com/dollarkillerx/fireworks/internal/app/backend/storage"
	"github.com/dollarkillerx/fireworks/internal/app/backend/task_pool"
	"github.com/dollarkillerx/fireworks/internal/conf"
	"github.com/gin-gonic/gin"
)

type Backend struct {
	app *gin.Engine
	db  storage.Interface

	taskPool *task_pool.TaskPool
}

func NewBackend(db storage.Interface) *Backend {
	return &Backend{
		db:       db,
		app:      gin.New(),
		taskPool: task_pool.New(),
	}
}

func (b *Backend) Run() error {
	if conf.GetBackendConfig().Debug {
		b.app.Use(gin.Logger())
	}
	b.router()

	return nil
}
