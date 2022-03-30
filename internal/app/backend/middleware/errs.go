package middleware

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"log"
)

// HttpRecover recover
func HttpRecover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover panic: ", err)
				utils.Return(ctx, errs.SystemError)
			}
		}()
	}
}

// SetBasicInformation 设置基础信息
func SetBasicInformation() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(enum.RequestID.String(), utils.GenerateID())
	}
}
