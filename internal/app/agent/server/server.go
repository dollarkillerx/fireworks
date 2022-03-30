package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/dollarkillerx/fireworks/internal/conf"
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/request"
	"github.com/dollarkillerx/processes"
	"github.com/dollarkillerx/urllib"
)

type AgentServer struct {
	conf *conf.AgentConfig
}

func NewAgentServer(conf *conf.AgentConfig) *AgentServer {
	return &AgentServer{
		conf: conf,
	}
}

func (a *AgentServer) Run() {
	go a.heartbeat()
	a.performTasks()
}

func (a *AgentServer) heartbeat() {
	for {
		time.Sleep(time.Second * 4)
		respCode, resp, err := urllib.Post("%s/agent/register").
			SetHeader("token", a.conf.Token).SetJsonObject(request.AddTask{
			AgentName:   a.conf.AgentName,
			AgentUrl:    a.conf.AgentIP,
			Workspace:   a.conf.Workspace,
			Description: a.conf.Description,
		}).Byte()
		if err != nil {
			log.Println(err)
			continue
		}

		if respCode != 200 {
			log.Println(string(resp))
		}
	}
}

func (a *AgentServer) performTasks() {
	for {
		time.Sleep(time.Second)
		err := a.performTaskCore()
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (a *AgentServer) performTaskCore() (err error) {
	defer func() {
		if err2 := recover(); err2 != nil {
			log.Println("Recover ...", err2)
		}
	}()

	var sub []models.Subtasks
	err = urllib.Post("/agent/recieve_task").SetHeader("token", a.conf.Token).
		SetJsonObject(request.AgentID{AgentID: a.conf.AgentName}).FromJson(&sub)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range sub {
		a.performTaskCoreItem(v)
	}

	return nil
}

func (a *AgentServer) performTaskCoreItem(sub models.Subtasks) {
	a.logs(sub.LogID, enum.TaskStatusRunning, enum.TaskStageBuild, "")
	log.Printf("task received task: %s log: %s \n", sub.ID, sub.TaskID)
	var instruction models.Instruction
	err := json.Unmarshal([]byte(sub.Instruction), &instruction)
	if err != nil {
		err = fmt.Errorf("Instruction 解析失败: " + err.Error())
		log.Println("Instruction 解析失败: " + err.Error())
		a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
		return
	}
	if len(instruction.Build) == 0 || len(instruction.Deploy) == 0 {
		err = fmt.Errorf("Instruction 解析失败: " + err.Error())
		log.Println("Instruction 解析失败: " + err.Error())
		a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
		return
	}

	exec := processes.NewExecLinux()
	_, err = exec.Exec("cd " + a.conf.Workspace)
	if err != nil {
		log.Println(err)
		a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
		return
	}

	pathStr, err := exec.Exec("pwd")
	if err != nil {
		log.Println(err)
		a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
		return
	}

	pathStr = path.Join(pathStr, fmt.Sprintf("%s_%s", sub.Name, sub.ID))
	err = os.MkdirAll(pathStr, 000777)
	if err != nil {
		log.Println(err)
		a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
		return
	}

	_, err = exec.Exec("cd " + pathStr)
	if err != nil {
		log.Println(err)
		a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
		return
	}

	switch sub.TaskType {
	case enum.TaskTypeDeploy:
		// build
		for _, v := range instruction.Build {
			r, err := exec.Exec(v)
			if err != nil {
				log.Println(err)
				a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
				return
			}
			log.Println(r)
		}
		// test
		for _, v := range instruction.Test {
			r, err := exec.Exec(v)
			if err != nil {
				log.Println(err)
				a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageTest, err.Error())
				return
			}
			log.Println(r)
		}
		// deploy
		for _, v := range instruction.Deploy {
			r, err := exec.Exec(v)
			if err != nil {
				log.Println(err)
				a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageDeploy, err.Error())
				return
			}
			log.Println(r)
		}

		a.logs(sub.LogID, enum.TaskStatusPassed, enum.TaskStageDeploy, err.Error())
	case enum.TaskTypeReboot:
		for _, v := range instruction.Reboot {
			r, err := exec.Exec(v)
			if err != nil {
				log.Println(err)
				a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageReboot, err.Error())
				return
			}
			log.Println(r)
		}
		a.logs(sub.LogID, enum.TaskStatusPassed, enum.TaskStageReboot, err.Error())
	case enum.TaskTypeStop:
		for _, v := range instruction.Stop {
			r, err := exec.Exec(v)
			if err != nil {
				log.Println(err)
				a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageStop, err.Error())
				return
			}
			log.Println(r)
		}
		a.logs(sub.LogID, enum.TaskStatusPassed, enum.TaskStageStop, err.Error())
	}
}

func (a *AgentServer) logs(logID string, taskStatus enum.TaskStatus, taskStage enum.TaskStage, logText string) {
	httpCode, resp, err := urllib.Post("/agent/task_log").SetHeader("token", a.conf.Token).
		SetJsonObject(request.TaskLogUpdate{
			LogID:      logID,
			TaskStage:  taskStage,
			TaskStatus: taskStatus,
			LogText:    logText,
		}).Byte()
	if err != nil {
		log.Println(err)
		return
	}

	if httpCode != 200 {
		log.Println(string(resp))
	}
}
