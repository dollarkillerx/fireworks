package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"log"
)

func (b *Backend) subtasks(ctx *gin.Context) {
	var task request.Task
	err := ctx.ShouldBindQuery(&task)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	tasks, err := b.db.GetSubtasks(task.TaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, tasks)
}

func (b *Backend) createSubtask(ctx *gin.Context) {
	var subtask request.CreateSubtask
	err := ctx.ShouldBindJSON(&subtask)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.CreateSubtasks(subtask.TaskID, subtask.Name, subtask.AgentID, subtask.Branch, subtask.Action, subtask.Instruction, subtask.Description)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}

func (b *Backend) disabledSubtask(ctx *gin.Context) {
	var subtask request.Subtask
	err := ctx.ShouldBindJSON(&subtask)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.DisabledSubtasks(subtask.SubtaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}

func (b *Backend) deleteSubtask(ctx *gin.Context) {
	var subtask request.Subtask
	err := ctx.ShouldBindJSON(&subtask)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.DelSubtasks(subtask.SubtaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}

func (b *Backend) updateSubtask(ctx *gin.Context) {
	var subtask request.UpdateSubtask
	err := ctx.ShouldBindJSON(&subtask)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.UpdateSubtasks(subtask.SubtaskID, subtask.Name, subtask.Instruction, subtask.Description)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}
