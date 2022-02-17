package server

import (
	"fmt"
	"log"
	"time"

	"github.com/dollarkillerx/fireworks/internal/conf"
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

	}

	return nil
}
