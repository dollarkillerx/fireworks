package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/pkg/git_models"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"fmt"
	"log"
	"strings"
)

func (b *Backend) gitlabTask(ctx *gin.Context) {
	token := ctx.GetHeader("X-Gitlab-Token")
	if token == "" {
		log.Println("gitlabTask error")
		utils.Return(ctx, errs.PleaseSignIn)
		return
	}

	var gitlabPayload git_models.GitlabWebHook
	err := ctx.ShouldBindJSON(&gitlabPayload)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.BadRequest)
		return
	}

	task, err := b.db.GetTaskByToken(token)
	if err != nil {
		log.Println("GetTaskByToken 不存在的任务: ", err.Error())
		utils.Return(ctx, errs.NewError("400002", "不存在的任务"))
		return
	}

	if task.Disabled {
		utils.Return(ctx, gin.H{})
		return
	}

	subtasks, err := b.db.GetSubtasks(task.ID)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SqlSystemError)
		return
	}

	for _, v := range subtasks {
		if v.Action == gitlabPayload.ObjectKind {
			if v.Disabled {
				continue
			}
			switch v.Action {
			case enum.TaskActionTag:
				if strings.Contains(gitlabPayload.Ref, v.Branch) {
					taskLog, err := b.db.CreateTaskLog(v.ID, enum.TaskTypeDeploy)
					if err != nil {
						log.Println(err)
						utils.Return(ctx, errs.SqlSystemError)
						return
					}
					v.LogID = taskLog
					v.TaskType = enum.TaskTypeDeploy
					b.taskPool.AddTask(v.AgentID, v)
				}
			case enum.TaskActionPush:
				if fmt.Sprintf("refs/heads/%s", v.Branch) == gitlabPayload.Ref {
					taskLog, err := b.db.CreateTaskLog(v.ID, enum.TaskTypeDeploy)
					if err != nil {
						log.Println(err)
						utils.Return(ctx, errs.SqlSystemError)
						return
					}
					v.LogID = taskLog
					v.TaskType = enum.TaskTypeDeploy
					b.taskPool.AddTask(v.AgentID, v)
				}
			case enum.TaskActionMerge:
				if gitlabPayload.ObjectAttributes == nil {
					continue
				}
				if v.Branch == gitlabPayload.ObjectAttributes.TargetBranch {
					taskLog, err := b.db.CreateTaskLog(v.ID, enum.TaskTypeDeploy)
					if err != nil {
						log.Println(err)
						utils.Return(ctx, errs.SqlSystemError)
						return
					}
					v.LogID = taskLog
					v.TaskType = enum.TaskTypeDeploy
					b.taskPool.AddTask(v.AgentID, v)
				}
			case enum.TaskActionPushOrMerge:
				if gitlabPayload.ObjectAttributes == nil {
					if fmt.Sprintf("refs/heads/%s", v.Branch) == gitlabPayload.Ref {
						taskLog, err := b.db.CreateTaskLog(v.ID, enum.TaskTypeDeploy)
						if err != nil {
							log.Println(err)
							utils.Return(ctx, errs.SqlSystemError)
							return
						}
						v.LogID = taskLog
						v.TaskType = enum.TaskTypeDeploy
						b.taskPool.AddTask(v.AgentID, v)
					}
					continue
				}
				if v.Branch == gitlabPayload.ObjectAttributes.TargetBranch {
					taskLog, err := b.db.CreateTaskLog(v.ID, enum.TaskTypeDeploy)
					if err != nil {
						log.Println(err)
						utils.Return(ctx, errs.SqlSystemError)
						return
					}
					v.LogID = taskLog
					v.TaskType = enum.TaskTypeDeploy
					b.taskPool.AddTask(v.AgentID, v)
				}
			}
		}
	}

	utils.Return(ctx, gin.H{})
}
