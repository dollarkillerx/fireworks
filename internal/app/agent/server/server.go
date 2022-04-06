package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dollarkillerx/fireworks/internal/utils"
	"log"
	"os"
	"path"
	"strings"
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
	lock *utils.RWLock
}

func NewAgentServer(conf *conf.AgentConfig) *AgentServer {
	return &AgentServer{
		conf: conf,
		lock: utils.NewRWLock(),
	}
}

func (a *AgentServer) Run() {
	go a.heartbeat()
	a.performTasks()
}

func (a *AgentServer) heartbeat() {
	for {
		respCode, resp, err := urllib.Post(fmt.Sprintf("%s/agent/register", a.conf.BackendAddr)).
			SetHeader("token", a.conf.Token).SetJsonObject(request.AddTask{
			AgentName:   a.conf.AgentName,
			AgentUrl:    a.conf.AgentIP,
			Workspace:   a.conf.Workspace,
			Description: a.conf.Description,
		}).Byte()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 4)
			continue
		}

		if respCode != 200 {
			log.Println(string(resp))
		}

		time.Sleep(time.Second * 4)
	}
}

func (a *AgentServer) performTasks() {
	for {
		err := a.performTaskCore()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
}

func (a *AgentServer) performTaskCore() (err error) {
	defer func() {
		if err2 := recover(); err2 != nil {
			log.Println("Recover ...", err2)
		}
	}()

	var sub []models.Subtasks
	err = urllib.Post(fmt.Sprintf("%s/agent/recieve_task", a.conf.BackendAddr)).SetHeader("token", a.conf.Token).
		SetJsonObject(request.AgentID{AgentID: a.conf.AgentName}).FromJson(&sub)
	if err != nil {
		log.Println(err)
		return
	}

	for i, v := range sub {
		a.logs(v.LogID, enum.TaskStatusRunning, enum.TaskStageBuild, "")
		log.Printf("task received task: %s log: %s \n", v.ID, v.TaskID)
		idx := i

		go a.performTaskCoreItem(sub[idx])
	}

	return nil
}

func (a *AgentServer) performTaskCoreItem(sub models.Subtasks) {
	lock := a.lock.Lock(sub.ID)
	defer lock.Unlock()

	if sub.GitAddr == "" {
		err := errors.New("Subtasks GitAddr is null " + sub.Name + "  " + sub.ID)
		log.Println(err)
		a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
		return
	}

	var instruction models.Instruction
	err := json.Unmarshal([]byte(sub.Instruction), &instruction)
	if err != nil {
		err = fmt.Errorf("Instruction 解析失败: " + err.Error() + " " + sub.Name + "  " + sub.ID)
		log.Println(err)
		a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
		return
	}
	if len(instruction.Build) == 0 || len(instruction.Deploy) == 0 {
		err = fmt.Errorf("Instruction 解析失败: " + err.Error() + " " + sub.Name + "  " + sub.ID)
		log.Println(err)
		a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
		return
	}

	exec := processes.NewExecLinuxGen("", "bash", time.Hour)
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

	ls, err := exec.Exec("ls -a")
	if err != nil {
		log.Println(err)
		a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
		return
	}

	if !strings.Contains(ls, ".git") {
		_, err := exec.Exec(fmt.Sprintf("git clone %s ./", sub.GitAddr))
		if err != nil {
			log.Println(err)
			a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
			return
		}
		_, err = exec.Exec(fmt.Sprintf("git checkout %s ", sub.Branch))
		if err != nil {
			log.Println(err)
			a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
			return
		}
	} else {
		_, err := exec.Exec("git checkout .")
		if err != nil {
			log.Println(err)
			a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
			return
		}

		_, err = exec.Exec("git pull")
		if err != nil {
			log.Println(err)
			a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
			return
		}
	}

	switch sub.TaskType {
	case enum.TaskTypeDeploy:
		var body string
		body += fmt.Sprintf("build =======\n")
		// build
		for _, v := range instruction.Build {
			r, err := exec.Exec(v)
			if err != nil {
				log.Println(err)
				a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageBuild, err.Error())
				return
			}
			body += fmt.Sprintf("%s\n", r)
		}

		body += fmt.Sprintf("test =======\n")
		// test
		for _, v := range instruction.Test {
			r, err := exec.Exec(v)
			if err != nil {
				log.Println(err)
				a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageTest, err.Error())
				return
			}
			log.Println(r)
			body += fmt.Sprintf("%s\n", r)
		}

		body += fmt.Sprintf("deploy =======\n")
		// deploy
		for _, v := range instruction.Deploy {
			r, err := exec.Exec(v)
			if err != nil {
				log.Println(err)
				a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageDeploy, err.Error())
				return
			}
			log.Println(r)
			body += fmt.Sprintf("%s\n", r)
		}

		a.logs(sub.LogID, enum.TaskStatusPassed, enum.TaskStageDeploy, body)
	case enum.TaskTypeReboot:
		var body string
		body += fmt.Sprintf("Reboot =======\n")
		for _, v := range instruction.Reboot {
			r, err := exec.Exec(v)
			if err != nil {
				log.Println(err)
				a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageReboot, err.Error())
				return
			}
			log.Println(r)
			body += fmt.Sprintf("%s\n", r)
		}
		a.logs(sub.LogID, enum.TaskStatusPassed, enum.TaskStageReboot, body)
	case enum.TaskTypeStop:
		var body string
		body += fmt.Sprintf("Stop =======\n")
		for _, v := range instruction.Stop {
			r, err := exec.Exec(v)
			if err != nil {
				log.Println(err)
				a.logs(sub.LogID, enum.TaskStatusFailed, enum.TaskStageStop, err.Error())
				return
			}
			log.Println(r)
			body += fmt.Sprintf("%s\n", r)
		}
		a.logs(sub.LogID, enum.TaskStatusPassed, enum.TaskStageStop, body)
	}
}

func (a *AgentServer) logs(logID string, taskStatus enum.TaskStatus, taskStage enum.TaskStage, logText string) {
	httpCode, resp, err := urllib.Post(fmt.Sprintf("%s/agent/task_log", a.conf.BackendAddr)).SetHeader("token", a.conf.Token).
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
