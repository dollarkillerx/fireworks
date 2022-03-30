package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"log"
	"time"
)

func (b *Backend) login(ctx *gin.Context) {
	var u request.WebLogin
	err := ctx.ShouldBind(&u)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	user, err := b.db.GetUser(u.Email)
	if err != nil {
		return
	}

	if u.Password != user.Password {
		utils.Return(ctx, errs.LoginFailed)
		return
	}

	// gen jwt
	token, err := utils.JWT.CreateToken(models.AuthToken{
		Email: user.Email,
		Name:  user.UserName,
		Role:  user.Role,
	}, int64(time.Hour*24*7))
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SystemError)
		return
	}

	utils.Return(ctx, gin.H{"token": token})
}

func (b *Backend) createUser(ctx *gin.Context) {
	var u request.CreateUser
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	model := utils.GetAuthModel(ctx)
	if model.Role != enum.Admin {
		utils.Return(ctx, errs.NewError("40001", "权限不足"))
		return
	}

	err = b.db.CreateUser(u.Email, u.Name, u.Password, enum.User)
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SystemError)
		return
	}

	utils.Return(ctx, gin.H{})
}
