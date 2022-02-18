package conf

import cfg "github.com/dollarkillerx/common/pkg/config"

type BackendConfig struct {
	ServerAddr string

	ListenAddr     string
	Debug          bool
	JWTToken       string
	PostgresConfig cfg.PostgresConfiguration
}

func InitBackendConfig(configName string, configPaths []string) (*BackendConfig, error) {
	var agentConfig BackendConfig
	err := cfg.InitConfiguration(configName, configPaths, &agentConfig)
	if err != nil {
		return nil, err
	}

	return &agentConfig, err
}
