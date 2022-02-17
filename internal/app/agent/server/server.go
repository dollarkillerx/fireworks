package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/dollarkillerx/fireworks/internal/conf"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/pkg/utils"
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

func (a *AgentServer) Run() error {
	// check
	code, bytes, err := urllib.Get(fmt.Sprintf("%s/info", a.conf.BackendAddr)).SetHeader("token", a.conf.Token).Byte()
	if err != nil {
		return err
	}
	if code != 200 {
		return fmt.Errorf("%s", bytes)
	}

	for {
		time.Sleep(time.Millisecond * 500)
		code, resps, err := urllib.Get(fmt.Sprintf("%s/info", a.conf.BackendAddr)).SetHeader("token", a.conf.Token).Byte()
		if err != nil {
			log.Println(err)
			continue
		}

		if code != 200 {
			log.Println(string(resps))
			continue
		}

		var agent models.Agent
		err = json.Unmarshal(resps, &agent)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return nil
}

func (a *AgentServer) deploy(agent models.Agent) {
	var err error
	os.MkdirAll(agent.TaskName, 00766)
	os.MkdirAll(path.Join(agent.TaskName, "files"), 00766)
	os.MkdirAll(path.Join(agent.TaskName, "images"), 00766)

	err = ioutil.WriteFile(path.Join(agent.TaskName, "docker-compose.yaml"), agent.DockerCompose, 00666)
	if err != nil {
		log.Println(err)
		return
	}

	err = a.download(path.Join(agent.TaskName, "files"), agent.Files)
	if err != nil {
		log.Println(err)
		return
	}
	err = a.download(path.Join(agent.TaskName, "images"), agent.DockerImages)
	if err != nil {
		log.Println(err)
		return
	}

}

func (a *AgentServer) download(path string, fp string) error {
	if fp == "" {
		return nil
	}
	code, resp, err := urllib.Get(fmt.Sprintf("%s/download/%s", a.conf.BackendAddr, fp)).SetHeader("token", a.conf.Token).Byte()
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(resp))
	}

	fileName := filepath.Join(path, "data.zip")
	err = ioutil.WriteFile(fileName, resp, 00766)
	if err != nil {
		return err
	}

	err = utils.DeCompress(fileName, path)
	if err != nil {
		return err
	}

	return nil
}
