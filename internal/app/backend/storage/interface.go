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
	GetTaskByToken(token string) (task *models.Task, err error)
	DisabledTask(taskID string) error
	UpdateTask(taskID string, name string, token string, description string) error
	DelTask(taskID string) error

	CreateSubtasks(taskID string, name string, agentID string, branch string, action enum.TaskAction, instruction string, description string) error
	GetSubtasks(taskID string) ([]models.Subtasks, error)
	GetSubtasksBySubtasksID(subtaskID string) (*models.Subtasks, error)
	DisabledSubtasks(subtaskID string) error
	DelSubtasks(subtaskID string) error
	UpdateSubtasks(subtaskID string, name string, instruction string, description string) error

	CreateTaskLog(subtasksID string, taskType enum.TaskType) (id string, err error)
	UpdateTaskLog(logID string, taskStatus enum.TaskStatus, taskStage enum.TaskStage, logText string) error
	GetTaskLog(subtasksID string) (logs []models.TaskLog, err error)
}
