package main

import (
	"flag"
	"log"
	"strings"

	"github.com/dollarkillerx/fireworks/internal/app/backend/server"
	"github.com/dollarkillerx/fireworks/internal/app/backend/storage/base"
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

	err := conf.InitBackendConfig(configFilename, strings.Split(configDirs, ","))
	if err != nil {
		log.Fatalln(err)
	}

	backend := server.NewBackend(base.NewBase())
	if err := backend.Run(); err != nil {
		log.Fatalln(err)
	}
}
