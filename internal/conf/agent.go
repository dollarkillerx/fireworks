package conf

import cfg "github.com/dollarkillerx/common/pkg/config"

type AgentConfig struct {
	BackendAddr string
	Token       string
}

func InitAgentConfig(configName string, configPaths []string) (*AgentConfig, error) {
	var agentConfig AgentConfig
	err := cfg.InitConfiguration(configName, configPaths, &agentConfig)
	if err != nil {
		return nil, err
	}

	return &agentConfig, err
}
