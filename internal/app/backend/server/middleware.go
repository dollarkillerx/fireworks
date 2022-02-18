package server

import (
	"encoding/json"
	"log"

	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/dollarkillerx/jwt"
	"github.com/gin-gonic/gin"
)

// setBasicInformation 设置基础信息
func setBasicInformation() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(enum.RequestID.String(), utils.GenerateID())
	}
}

func authToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("token")
		if tokenStr == "" {
			utils.Return(ctx, errs.PleaseSignIn)
			return
		}

		token, err := jwt.TokenFormatString(tokenStr)
		if err != nil {
			utils.Return(ctx, errs.PleaseSignIn)
			return
		}

		err = utils.JWT.VerificationSignature(token)
		if err != nil {
			utils.Return(ctx, errs.PleaseSignIn)
			return
		}

		marshal, err := json.Marshal(token.Payload.Payload)
		if err != nil {
			log.Println(err)
			utils.Return(ctx, errs.PleaseSignIn)
			return
		}

		var authModel request.AuthModel
		err = json.Unmarshal(marshal, &authModel)
		if err != nil {
			log.Println(err)
			utils.Return(ctx, errs.PleaseSignIn)
			return
		}

		err = utils.Validate.Struct(&authModel)
		if err != nil {
			log.Println(err)
			utils.Return(ctx, errs.PleaseSignIn)
			return
		}

		ctx.Set(enum.AuthModel.String(), authModel)

		ctx.Next()
	}
}
