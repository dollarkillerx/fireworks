package server

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/errs"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"github.com/gin-gonic/gin"

	"log"
)

func (b *Backend) agents(ctx *gin.Context) {
	agent, err := b.db.ListAgent()
	if err != nil {
		log.Println(err)
		utils.Return(ctx, errs.SystemError)
		return
	}

	utils.Return(ctx, agent)
}
