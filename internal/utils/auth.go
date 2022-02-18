package utils

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/gin-gonic/gin"
)

// GetAuthModel GetAuthModel
func GetAuthModel(ctx *gin.Context) request.AuthModel {
	get, exists := ctx.Get(enum.AuthModel.String())
	if !exists {
		panic("what fuck GetAuthModel is not exists")
	}

	model, ok := get.(request.AuthModel)
	if !ok {
		panic("what fuck GetAuthModel is not exists 2")
	}

	return model
}
