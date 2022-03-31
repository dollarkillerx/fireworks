package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"log"
)

func (b *Backend) taskLogs(ctx *gin.Context) {
	var logID request.TaskLogID
	err := ctx.ShouldBindQuery(&logID)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	if logID.TaskID == "undefined" {
		logID.TaskID = ""
	}
	if logID.SubtaskID == "undefined" {
		logID.SubtaskID = ""
	}

	taskLogs, err := b.db.GetTaskLog(logID.TaskID, logID.SubtaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, taskLogs)
}
