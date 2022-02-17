package models

type Agent struct {
	TaskID        string `json:"task_id"`
	TaskName      string `json:"task_name"`
	DockerCompose []byte `json:"docker_compose"`
	DockerImages  string `json:"docker_images"` // docker 文件地址
	Files         string `json:"files"`         // file 文件地址
}
