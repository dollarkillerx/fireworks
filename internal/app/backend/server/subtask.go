package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"encoding/json"
	"log"
)

func (b *Backend) subtasks(ctx *gin.Context) {
	var task request.Task
	err := ctx.ShouldBindQuery(&task)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	agent, err := b.db.ListAgent()
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	tasks, err := b.db.GetSubtasks(task.TaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	for i, v := range tasks {
		for _, vv := range agent {
			if v.AgentID == vv.ID {
				tasks[i].AgentName = vv.AgentName
			}
		}
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

	var instruction models.Instruction
	err = json.Unmarshal([]byte(subtask.Instruction), &instruction)
	if err != nil {
		utils.Return(ctx, errs.NewError("400003", "Instruction 解析失败: "+err.Error()))
		return
	}
	if len(instruction.Build) == 0 || len(instruction.Deploy) == 0 {
		utils.Return(ctx, errs.NewError("400003", "Instruction 填写错误"))
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

	var instruction models.Instruction
	err = json.Unmarshal([]byte(subtask.Instruction), &instruction)
	if err != nil {
		utils.Return(ctx, errs.NewError("400003", "Instruction 解析失败: "+err.Error()))
		return
	}
	if len(instruction.Build) == 0 || len(instruction.Deploy) == 0 {
		utils.Return(ctx, errs.NewError("400003", "Instruction 解析失败: "+err.Error()))
		return
	}

	err = b.db.UpdateSubtasks(subtask.SubtaskID, subtask.Instruction, subtask.Description)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}

func (b *Backend) rebootSubtask(ctx *gin.Context) {
	var subtask request.Subtask
	err := ctx.ShouldBindJSON(&subtask)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	id, err := b.db.GetSubtasksBySubtasksID(subtask.SubtaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.BadRequest)
		return
	}

	taskLog, err := b.db.CreateTaskLog(subtask.SubtaskID, enum.TaskTypeReboot)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	id.LogID = taskLog
	id.TaskType = enum.TaskTypeReboot
	b.taskPool.AddTask(id.AgentID, *id)

	utils.Return(ctx, gin.H{})
}

func (b *Backend) stopSubtask(ctx *gin.Context) {
	var subtask request.Subtask
	err := ctx.ShouldBindJSON(&subtask)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	id, err := b.db.GetSubtasksBySubtasksID(subtask.SubtaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.BadRequest)
		return
	}

	taskLog, err := b.db.CreateTaskLog(subtask.SubtaskID, enum.TaskTypeStop)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	id.LogID = taskLog
	id.TaskType = enum.TaskTypeStop
	b.taskPool.AddTask(id.AgentID, *id)

	utils.Return(ctx, gin.H{})
}
