package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/pkg/git_models"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func (b *Backend) webHookTask(ctx *gin.Context) {
	token := ctx.GetHeader("X-Gitlab-Token")
	if token != "" {
		b.gitlabTask(ctx, token)
		return
	}

	log.Println("未知请求")
	utils.Return(ctx, errs.PleaseSignIn)
	return
}

func (b *Backend) gitlabTask(ctx *gin.Context, token string) {
	var gitlabPayload git_models.GitlabWebHook
	err := ctx.ShouldBindJSON(&gitlabPayload)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.BadRequest)
		return
	}

	fmt.Println("=========gitlab payload: ")
	marshal, err := json.Marshal(gitlabPayload)
	if err != nil {
		panic(err)
	}
	log.Println(string(marshal))
	fmt.Println("=========token: ")
	fmt.Println(token)

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

	gitSsh := gitlabPayload.Project.GitSshUrl

	for _, v := range subtasks {
		if v.Action == gitlabPayload.ObjectKind || (v.Action == enum.TaskActionPushOrMerge && gitlabPayload.ObjectKind == "push") || (v.Action == enum.TaskActionPushOrMerge && gitlabPayload.ObjectKind == "merge_request") {
			if v.Disabled {
				continue
			}
			switch v.Action {
			case enum.TaskActionTag:
				if strings.Contains(gitlabPayload.Ref, v.Branch) {
					_, err := b.db.CreateTaskLog(v.ID, gitSsh, enum.TaskTypeDeploy)
					if err != nil {
						log.Println(err)
						utils.Return(ctx, errs.SqlSystemError)
						return
					}
				}
			case enum.TaskActionPush:
				if fmt.Sprintf("refs/heads/%s", v.Branch) == gitlabPayload.Ref {
					_, err := b.db.CreateTaskLog(v.ID, gitSsh, enum.TaskTypeDeploy)
					if err != nil {
						log.Println(err)
						utils.Return(ctx, errs.SqlSystemError)
						return
					}
				}
			case enum.TaskActionMerge:
				if gitlabPayload.ObjectAttributes == nil {
					continue
				}
				if v.Branch == gitlabPayload.ObjectAttributes.TargetBranch {
					_, err := b.db.CreateTaskLog(v.ID, gitSsh, enum.TaskTypeDeploy)
					if err != nil {
						log.Println(err)
						utils.Return(ctx, errs.SqlSystemError)
						return
					}
				}
			case enum.TaskActionPushOrMerge:
				if gitlabPayload.ObjectAttributes == nil {
					if fmt.Sprintf("refs/heads/%s", v.Branch) == gitlabPayload.Ref {
						_, err := b.db.CreateTaskLog(v.ID, gitSsh, enum.TaskTypeDeploy)
						if err != nil {
							log.Println(err)
							utils.Return(ctx, errs.SqlSystemError)
							return
						}
					} else {
						log.Println(fmt.Sprintf("refs/heads/%s", v.Branch))
						log.Println(fmt.Sprintf(gitlabPayload.Ref))
					}
					continue
				}
				if v.Branch == gitlabPayload.ObjectAttributes.TargetBranch {
					_, err := b.db.CreateTaskLog(v.ID, gitSsh, enum.TaskTypeDeploy)
					if err != nil {
						log.Println(err)
						utils.Return(ctx, errs.SqlSystemError)
						return
					}
				}
			}
		}
	}

	utils.Return(ctx, gin.H{})
}
