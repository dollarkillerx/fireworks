package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"time"
)

func (b *Backend) login(ctx *gin.Context) {
	var u request.WebLogin
	err := ctx.ShouldBind(&u)
	if err != nil {
		utils.Return(ctx, errs.BadRequest)
		return
	}

	var user models.User
	err = b.db.Model(&models.User{}).Where("email = ?", u.Email).First(&user).Error
	if err != nil {
		utils.Return(ctx, errs.LoginFailed)
		return
	}

	if u.Password != user.Password {
		utils.Return(ctx, errs.LoginFailed)
		return
	}

	// gen jwt
	token, err := utils.JWT.CreateToken(map[string]string{
		"email": user.Email,
		"name":  user.UserName,
		"role":  string(user.Role),
	}, time.Now().Add(time.Hour*24*6).Unix())
	if err != nil {
		utils.Return(ctx, errs.SystemError)
		return
	}

	utils.Return(ctx, gin.H{"token": token})
}

func (b *Backend) createUser(ctx *gin.Context) {
	b.db.Model(&models.User{})
}
