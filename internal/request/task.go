package request

import "github.com/dollarkillerx/fireworks/internal/pkg/enum"

type CreateTask struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
}

type Task struct {
	TaskID string `json:"task_id" form:"task_id" binding:"required"`
}

type Subtask struct {
	SubtaskID string `json:"subtask_id" form:"subtask_id" binding:"required"`
}

type TaskLogID struct {
	TaskID    string `json:"task_id" form:"task_id"`
	SubtaskID string `json:"subtask_id" form:"subtask_id"`
}

type CreateSubtask struct {
	TaskID      string          `json:"task_id" binding:"required"`
	Name        string          `json:"name" binding:"required"`
	AgentID     string          `json:"agent_id" binding:"required"`
	Branch      string          `json:"branch" binding:"required"`
	Action      enum.TaskAction `json:"action" binding:"required"`
	Instruction string          `json:"instruction" binding:"required"` // 指令
	Description string          `json:"description"`
}

type UpdateSubtask struct {
	SubtaskID   string `json:"subtask_id" form:"subtask_id" binding:"required"`
	Instruction string `json:"instruction" binding:"required"` // 指令
	Description string `json:"description"`
}

type CreateConfiguration struct {
	Subtask  string `json:"subtask" binding:"required"`
	Filename string `json:"filename" binding:"required"`
	Body     string `json:"body" binding:"required"`
}

type ModifyConfiguration struct {
	ConfigurationID string `json:"configuration_id" binding:"required"`
	Filename        string `json:"filename"`
	Body            string `json:"body"`
}

type ConfigurationID struct {
	ConfigurationID string `json:"configuration_id" form:"configuration_id"  binding:"required"`
}

type ConfigurationToken struct {
	ConfigurationToken string `json:"configuration_token" form:"configuration_token"  binding:"required"`
}
