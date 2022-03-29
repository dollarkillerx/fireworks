package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.POST("/task", func(ctx *gin.Context) {
		header := ctx.GetHeader("X-Gitlab-Token HTTP")
		fmt.Println(header)

		defer ctx.Request.Body.Close()
		file, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(file))
	})

	fmt.Println("run 0.0.0.0:8754")
	err := app.Run("0.0.0.0:8754")
	if err != nil {
		panic(err)
	}
}
