package server

import (
	"github.com/dollarkillerx/fireworks/internal/conf"
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"log"
)

func (b *Backend) agents(ctx *gin.Context) {
	agent, err := b.db.ListAgent()
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SystemError)
		return
	}

	utils.Return(ctx, agent)
}

func (b *Backend) agentAuth(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != conf.GetBackendConfig().AgentToken {
		utils.Return(ctx, errs.PleaseSignIn)
		return
	}
}

func (b *Backend) registerAgent(ctx *gin.Context) {
	var add request.AddTask
	err := ctx.ShouldBindJSON(&add)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.AddAgent(add.AgentName, add.AgentUrl, add.Workspace, add.Description)
	if err != nil {
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}

func (b *Backend) recieveTask(ctx *gin.Context) {
	var add request.AgentID
	err := ctx.ShouldBindJSON(&add)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	tasks, err := b.db.RecieveTaskByLog(add.AgentID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.BadRequest)
		return
	}

	ctx.JSON(200, tasks)
}

func (b *Backend) taskLog(ctx *gin.Context) {
	var add request.TaskLogUpdate
	err := ctx.ShouldBindJSON(&add)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.UpdateTaskLog(add.LogID, add.TaskStatus, add.TaskStage, add.LogText)
	if err != nil {
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}
