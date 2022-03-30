package server

import (
	"github.com/dollarkillerx/fireworks/internal/app/backend/middleware"
	"github.com/gin-gonic/gin"
)

func (b *Backend) router() {
	b.app.Use(middleware.HttpRecover())
	b.app.Use(middleware.Cors())
	b.app.Use(middleware.SetBasicInformation())
	b.app.NoRoute(func(context *gin.Context) {
		context.Redirect(302, "/")
	})

	b.app.POST("/api/web/login", b.login)
	v1Api := b.app.Group("/api/web/v1", middleware.Auth())
	{
		// user
		v1Api.POST("create_user", b.createUser)

		// agent
		v1Api.GET("agents", b.agents)

		// task

		task := v1Api.Group("/task")
		{
			task.GET("/", b.tasks)
			task.POST("create", b.createTask)
			task.POST("disabled", b.disabledTask)
			task.POST("delete", b.deleteTask)
		}

		// subtask

		subtask := v1Api.Group("/subtask")
		{
			subtask.GET("/", b.subtasks)
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

	b.app.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		//content, err := ioutil.ReadFile("dist/index.html")
		//if err != nil {
		//	c.Writer.WriteHeader(404)
		//	c.Writer.WriteString("Not Found")
		//	return
		//}
		//_, _ = c.Writer.Write(content)
		//c.Writer.Header().Add("Accept", "text/html")
		//c.Writer.Flush()
	})
}
