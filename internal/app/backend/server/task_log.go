package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"log"
)

func (b *Backend) taskLogs(ctx *gin.Context) {
	var subtask request.Subtask
	err := ctx.ShouldBindJSON(&subtask)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	taskLogs, err := b.db.GetTaskLog(subtask.SubtaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, taskLogs)
}
