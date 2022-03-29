package storage

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"gorm.io/gorm"
)

type Interface interface {
	GetDB() *gorm.DB

	user
	agent
	task
}

type user interface {
	CreateUser(email string, name string, password string, role enum.Role) error
	GetUser(email string) (*models.User, error)
}

type agent interface {
	AddAgent(agentName string, agentUrl string, workspace string, description string) error
	ListAgent() ([]models.Agent, error)
}

type task interface {
	CreateTask(name string, description string) error
	GetTasks() ([]models.Task, error)
	UpdateTask(taskID string, name string, token string, description string) error
	DelTask(taskID string) error

	CreateSubtasks(taskID string, name string, agentID string, branch string, action enum.TaskAction, instruction string, description string) error
	GetSubtasks(taskID string) ([]models.Subtasks, error)
	DisabledSubtasks(subtasksID string) error
	DelSubtasks(subtasksID string) error

	CreateTaskLog(subtasksID string) (id string, err error)
	UpdateTaskLog(logID string, taskStatus enum.TaskStatus, taskStage enum.TaskStage, logText string) error
	GetTaskLog(subtasksID string) (logs []models.TaskLog, err error)
}
