package server

import (
	"github.com/dollarkillerx/fireworks/internal/app/backend/middleware"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (b *Backend) router() {
	b.app.Use(middleware.HttpRecover())
	b.app.Use(middleware.Cors())
	b.app.Use(middleware.SetBasicInformation())
	b.app.NoRoute(func(context *gin.Context) {
		context.Redirect(302, "/")
	})

	b.app.POST("/api/web/v1/login", b.login)
	v1Api := b.app.Group("/api/web/v1", middleware.Auth())
	{
		// user
		v1Api.POST("create_user", b.createUser)

		// agent
		v1Api.GET("agents", b.agents)

		// task

		task := v1Api.Group("/task")
		{
			task.GET("index", b.tasks)
			task.POST("create", b.createTask)
			task.POST("disabled", b.disabledTask)
			task.POST("delete", b.deleteTask)
		}

		// subtask

		subtask := v1Api.Group("/subtask")
		{
			subtask.GET("index", b.subtasks)
			subtask.POST("create", b.createSubtask)
			subtask.POST("disabled", b.disabledSubtask)
			subtask.POST("delete", b.deleteSubtask)
			subtask.POST("update", b.updateSubtask)
			subtask.POST("reboot", b.rebootSubtask)
			subtask.POST("stop", b.stopSubtask)
		}

		v1Api.GET("task_logs", b.taskLogs)
	}

	webhook := b.app.Group("/webhook")
	{
		webhook.POST("/task/gitlab", b.gitlabTask)
	}

	agent := b.app.Group("/agent", b.agentAuth)
	{
		agent.POST("/register", b.registerAgent)
		agent.POST("/recieve_task", b.recieveTask)
		agent.POST("/task_log", b.taskLog)
	}

	b.app.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		content, err := ioutil.ReadFile("dist/index.html")
		if err != nil {
			c.Writer.WriteHeader(404)
			c.Writer.WriteString("Not Found")
			return
		}
		_, _ = c.Writer.Write(content)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Flush()
	})

	b.app.StaticFile("go-logo-blue.svg", "./dist/go-logo-blue.svg")
	b.app.StaticFile("cicd-pipeline.png", "./dist/cicd-pipeline.png")
	b.app.StaticFS("/static/css", http.Dir("./dist/static/css"))
	b.app.StaticFS("/static/fonts", http.Dir("./dist/static/fonts"))
	b.app.StaticFS("/static/img", http.Dir("./dist/static/img"))
	b.app.StaticFS("/static/js", http.Dir("./dist/static/js"))
}
