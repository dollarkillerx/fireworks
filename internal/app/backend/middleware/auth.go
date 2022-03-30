package middleware

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/dollarkillerx/jwt"
	"github.com/gin-gonic/gin"

	"log"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenS := ctx.GetHeader("token")
		if tokenS == "" {
			utils.Return(ctx, errs.PleaseSignIn)
			return
		}

		token, err := jwt.TokenFormatString(tokenS)
		if err != nil {
			log.Println(err)
			utils.Return(ctx, errs.PleaseSignIn)
			return
		}

		var tk models.AuthToken
		err = token.Payload.Unmarshal(&tk)
		if err != nil {
			log.Println(err)
			utils.Return(ctx, errs.PleaseSignIn)
			return
		}

		ctx.Set(enum.AuthModel.String(), tk)
	}
}
