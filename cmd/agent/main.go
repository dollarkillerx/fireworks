package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/dollarkillerx/fireworks/internal/app/agent/server"
	"github.com/dollarkillerx/fireworks/internal/conf"
)

var configFilename string
var configDirs string

func init() {
	const (
		defaultConfigFilename = "config"
		configUsage           = "Name of the config file, without extension"
		defaultConfigDirs     = "./,./configs/"
		configDirUsage        = "Directories to search for config file, separated by ','"
	)
	flag.StringVar(&configFilename, "c", defaultConfigFilename, configUsage)
	flag.StringVar(&configDirs, "cPath", defaultConfigDirs, configDirUsage)
	flag.Parse()
}

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	conf, err := conf.InitAgentConfig(configFilename, strings.Split(configDirs, ","))
	if err != nil {
		log.Fatalln(err)
	}

	agentServer := server.NewAgentServer(conf)
	fmt.Println("Fireworks Agent Run ...")
	if err := agentServer.Run(); err != nil {
		panic(err)
	}
}
