package models

type Agent struct {
	BaseModel
	AgentName   string `gorm:"type:varchar(600);index" json:"agent_name"`
	AgentUrl    string `gorm:"type:varchar(600);index" json:"agent_url"`
	Workspace   string `gorm:"type:text" json:"workspace"`
	Description string `gorm:"type:text" json:"description"`
	Expired     int64  `gorm:"index" json:"expired"`
	Live        bool   `json:"live" gorm:"-"` // 是否存活
}
