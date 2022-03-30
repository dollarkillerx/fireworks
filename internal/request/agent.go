package request

import "github.com/dollarkillerx/fireworks/internal/pkg/enum"

type AddTask struct {
	AgentName   string `json:"agent_name" binding:"required"`
	AgentUrl    string `json:"agent_url" binding:"required"`
	Workspace   string `json:"workspace" binding:"required"`
	Description string `json:"description"`
}

type AgentID struct {
	AgentID string `json:"agent_id" binding:"required"`
}

type TaskLogUpdate struct {
	LogID      string          `json:"log_id"`
	TaskStatus enum.TaskStatus `json:"task_status"`
	TaskStage  enum.TaskStage  `json:"task_stage"`
	LogText    string          `json:"log_text"`
}
