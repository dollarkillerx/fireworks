package models

import "github.com/dollarkillerx/fireworks/internal/pkg/enum"

type Agent struct {
	BaseModel
	AgentName   string `gorm:"type:varchar(600);index" json:"agent_name"`
	AgentUrl    string `gorm:"type:varchar(600);index" json:"agent_url"`
	Workspace   string `gorm:"type:text" json:"workspace"`
	Description string `gorm:"type:text" json:"description"`
}

// TaskLog 任务日志
type TaskLog struct {
	BaseModel
	TaskID     string          `gorm:"type:varchar(600);index" json:"task_id"`
	TaskStatus enum.TaskStatus `gorm:"type:varchar(600);index" json:"task_status"`
	TaskStage  enum.TaskStage  `gorm:"type:varchar(600);index" json:"task_stage"`
	Log        string          `gorm:"type:text"`
}

// Task 具体任务
type Task struct {
	BaseModel
	Name        string `gorm:"type:varchar(600);index" json:"name"`  // task name
	Token       string `gorm:"type:varchar(600);index" json:"token"` // token
	Description string `gorm:"type:text" json:"description"`
}

// Subtasks 子任务
type Subtasks struct {
	BaseModel
	Name        string `gorm:"type:varchar(600);index" json:"name"`     // subtasks name
	TaskID      string `gorm:"type:varchar(600);index" json:"task_id"`  // 任务ID
	AgentID     string `gorm:"type:varchar(600);index" json:"agent_id"` // agentID
	Branch      string `gorm:"type:varchar(600);index" json:"branch"`   // 分支
	Instruction string `gorm:"type:text" json:"instruction"`            // 指令
	Description string `gorm:"type:text" json:"description"`
}
