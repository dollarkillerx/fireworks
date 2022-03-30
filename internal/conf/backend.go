package conf

import cfg "github.com/dollarkillerx/common/pkg/config"

type backendConfig struct {
	ServerAddr string

	ListenAddr         string
	Debug              bool
	JWTToken           string
	BasicAdministrator Account
	PostgresConfig     cfg.PostgresConfiguration
}

type Account struct {
	Email    string
	Name     string
	Password string
}

var backendConfigInternal *backendConfig

func InitBackendConfig(configName string, configPaths []string) error {
	var agentConfig backendConfig
	err := cfg.InitConfiguration(configName, configPaths, &agentConfig)
	if err != nil {
		return err
	}

	backendConfigInternal = &agentConfig

	return err
}

func GetBackendConfig() *backendConfig {
	return backendConfigInternal
}
