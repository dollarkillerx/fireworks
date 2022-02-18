package utils

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/gin-gonic/gin"
)

// UniversalReturn 通用返回
type UniversalReturn struct {
	RequestID string      `json:"request_id,omitempty"` // 请求ID
	Code      string      `json:"code"`                 // 状态码  0 正常
	Message   string      `json:"message,omitempty"`    // 基础消息
	Data      interface{} `json:"data,omitempty"`       // 数据
}

func Return(ctx *gin.Context, data interface{}) {
	switch r := data.(type) {
	case *errs.Error:
		if r.HttpCode == 0 {
			r.HttpCode = 200
		}

		requestID, _ := ctx.Get(enum.RequestID.String())
		ctx.JSON(r.HttpCode, UniversalReturn{
			RequestID: requestID.(string),
			Code:      r.Code,
			Message:   r.Message,
		})
	case errs.Error:
		if r.HttpCode == 0 {
			r.HttpCode = 200
		}

		requestID, _ := ctx.Get(enum.RequestID.String())
		ctx.JSON(r.HttpCode, UniversalReturn{
			RequestID: requestID.(string),
			Code:      r.Code,
			Message:   r.Message,
		})
	default:
		ctx.JSON(200, UniversalReturn{
			Code: "0",
			Data: data,
		})
	}

	ctx.Abort()
}
