package models

type Agent struct {
	TaskID        string   `json:"task_id"`
	TaskName      string   `json:"task_name"`
	DockerCompose []byte   `json:"docker_compose"`
	DockerAddr    []string `json:"docker_addr"` // docker 文件地址
	Files         []string `json:"files"`       // file 文件地址
}
