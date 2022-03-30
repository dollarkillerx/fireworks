package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"log"
)

func (b *Backend) tasks(ctx *gin.Context) {
	tasks, err := b.db.GetTasks()
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, tasks)
}

func (b *Backend) createTask(ctx *gin.Context) {
	var task request.CreateTask
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.CreateTask(task.Name, task.Description)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.NewError("50001", err.Error()))
		return
	}

	utils.Return(ctx, gin.H{})
}

func (b *Backend) disabledTask(ctx *gin.Context) {
	var task request.Task
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.DisabledTask(task.TaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}

func (b *Backend) deleteTask(ctx *gin.Context) {
	var task request.Task
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.DelTask(task.TaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}
