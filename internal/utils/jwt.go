package utils

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/jwt"
	"github.com/gin-gonic/gin"
)

var JWT *jwt.JWT

func InitJWT() {
	JWT = jwt.NewJwt(config.GetConfig().JWTToken)
}

// GetAuthModel GetAuthModel
func GetAuthModel(ctx *gin.Context) models.AuthToken {
	get, exists := ctx.Get(enum.AuthModel.String())
	if !exists {
		panic("what fuck AuthToken is not exists")
	}

	model, ok := get.(models.AuthToken)
	if !ok {
		panic("what fuck AuthToken is not exists 2")
	}

	return model
}
