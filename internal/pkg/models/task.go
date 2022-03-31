package models

import "github.com/dollarkillerx/fireworks/internal/pkg/enum"

// TaskLog 任务日志
type TaskLog struct {
	BaseModel
	TaskID      string          `gorm:"type:varchar(600);index" json:"task_id"`
	TaskName    string          `gorm:"type:varchar(600)" json:"task_name"`
	SubtasksID  string          `gorm:"type:varchar(600);index" json:"subtasks_id"`
	SubtaskName string          `gorm:"type:varchar(600)" json:"subtask_name"`
	TaskType    enum.TaskType   `gorm:"type:varchar(600);index" json:"task_type"`
	TaskStatus  enum.TaskStatus `gorm:"type:varchar(600);index" json:"task_status"`
	TaskStage   enum.TaskStage  `gorm:"type:varchar(600);index" json:"task_stage"`
	Log         string          `gorm:"type:text"`

	DisplayTime string `gorm:"-" json:"display_time"`
}

// Task 具体任务
type Task struct {
	BaseModel
	Name        string `gorm:"type:varchar(600);unique" json:"name"` // task name
	Token       string `gorm:"type:varchar(600);index" json:"token"` // token
	Description string `gorm:"type:text" json:"description"`
	Disabled    bool   `gorm:"index" json:"disabled"` //　禁用
}

// Subtasks 子任务
type Subtasks struct {
	BaseModel
	Name        string          `gorm:"type:varchar(600);index" json:"name"`     // subtasks name
	TaskID      string          `gorm:"type:varchar(600);index" json:"task_id"`  // 任务ID
	AgentID     string          `gorm:"type:varchar(600);index" json:"agent_id"` // agentID
	Branch      string          `gorm:"type:varchar(600);index" json:"branch"`   // 分支
	Action      enum.TaskAction `gorm:"type:varchar(600);index" json:"action"`   // action
	Instruction string          `gorm:"type:text" json:"instruction"`            // 指令
	Description string          `gorm:"type:text" json:"description"`
	Payload     string          `gorm:"type:text" json:"payload"` // git payload
	Disabled    bool            `gorm:"index" json:"disabled"`    //　禁用

	LogID     string        `json:"log_id" gorm:"-"`     // 其他字段使用
	TaskType  enum.TaskType `gorm:"-" json:"task_type"`  // 其他字段使用
	AgentName string        `gorm:"-" json:"agent_name"` // 其他字段使用
}

type Instruction struct {
	Build  []string `json:"build"`
	Test   []string `json:"test"`
	Deploy []string `json:"deploy"`
	Reboot []string `json:"reboot"`
	Stop   []string `json:"stop"`
	Delete []string `json:"delete"`
}
