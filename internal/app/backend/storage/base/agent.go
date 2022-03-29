package base

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/utils"

	"log"
	"time"
)

func (b *Base) AddAgent(agentName string, agentUrl string, workspace string, description string) error {
	var count int64
	err := b.db.Model(&models.Agent{}).Where("agent_name = ?", agentName).Count(&count).Error
	if err != nil {
		log.Println(err)
		return err
	}

	if count == 0 {
		err := b.db.Model(&models.Agent{}).Create(&models.Agent{
			BaseModel: models.BaseModel{
				ID: utils.GenerateID(),
			},
			AgentName:   agentName,
			AgentUrl:    agentUrl,
			Workspace:   workspace,
			Description: description,
			Expired:     time.Now().Add(time.Second * 6).Unix(),
		}).Error
		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	}

	err = b.db.Model(&models.Agent{}).Where("agent_name = ?", agentName).
		Updates(&models.Agent{
			AgentUrl:    agentUrl,
			Workspace:   workspace,
			Description: description,
			Expired:     time.Now().Add(time.Second * 6).Unix(),
		}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) ListAgent() (agent []models.Agent, err error) {
	err = b.db.Model(&models.Agent{}).Find(&agent).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for i := range agent {
		if agent[i].Expired > time.Now().Unix() {
			agent[i].Live = true
		}
	}

	return
}
