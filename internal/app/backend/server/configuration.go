package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/fireworks/internal/response"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"log"
)

func (b *Backend) configuration(ctx *gin.Context) {
	var subtask request.Subtask
	err := ctx.ShouldBindQuery(&subtask)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	configurations, err := b.db.Configurations(subtask.SubtaskID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, configurations)
}

func (b *Backend) configurationItem(ctx *gin.Context) {
	var req request.ConfigurationID
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.BadRequest)
		return
	}

	id, err := b.db.GetConfigurationByID(req.ConfigurationID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.BadRequest)
		return
	}

	utils.Return(ctx, id)
}

func (b *Backend) createConfiguration(ctx *gin.Context) {
	var req request.CreateConfiguration
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.CreateConfiguration(req.Subtask, req.Filename, req.Body)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}

func (b *Backend) modifyConfiguration(ctx *gin.Context) {
	var req request.ModifyConfiguration
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.ModifyConfiguration(req.ConfigurationID, req.Filename, req.Body)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}

func (b *Backend) deleteConfiguration(ctx *gin.Context) {
	var req request.ConfigurationID
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	err = b.db.DeleteConfiguration(req.ConfigurationID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}

func (b *Backend) configurations(ctx *gin.Context) {
	var req request.ConfigurationToken
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		log.Println()
		utils.Return(ctx, errs.BadRequest)
		return
	}

	subtasks, err := b.db.GetSubtasksByToken(req.ConfigurationToken)
	if err != nil {
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	cfgs, err := b.db.GetConfigurationBySubtaskID(subtasks.ID)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	var resp response.Configurations
	for _, v := range cfgs {
		resp.Configs = append(resp.Configs, response.ConfigurationItem{
			Filename: v.Filename,
			Body:     v.Body,
		})
	}

	ctx.JSON(200, resp)
}
