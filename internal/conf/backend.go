package conf

import cfg "github.com/dollarkillerx/common/pkg/config"

type BackendConfig struct {
	ListenAddr string
	Key        string
}

func InitBackendConfig(configName string, configPaths []string) (*BackendConfig, error) {
	var agentConfig BackendConfig
	err := cfg.InitConfiguration(configName, configPaths, &agentConfig)
	if err != nil {
		return nil, err
	}

	return &agentConfig, err
}
