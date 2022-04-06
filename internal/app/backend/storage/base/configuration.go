package base

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/utils"

	"log"
)

func (b *Base) Configurations(subtask string) (cfgs []models.Configuration, err error) {
	err = b.db.Model(&models.Configuration{}).Where("subtask_id = ?", subtask).Find(&cfgs).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}

func (b *Base) CreateConfiguration(subtask string, filename string, body string) error {
	err := b.db.Model(&models.Configuration{}).Create(&models.Configuration{
		BaseModel: models.BaseModel{ID: utils.GenerateID()},
		SubtaskID: subtask,
		Filename:  filename,
		Body:      body,
	}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) GetConfigurationByID(id string) (*models.Configuration, error) {
	var conf models.Configuration
	err := b.db.Model(&models.Configuration{}).Where("id = ?", id).First(&conf).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &conf, nil
}

func (b *Base) ModifyConfiguration(configurationID string, filename string, body string) error {
	err := b.db.Model(&models.Configuration{}).Where("id = ?", configurationID).Updates(&models.Configuration{
		Filename: filename,
		Body:     body,
	}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) DeleteConfiguration(configurationID string) error {
	err := b.db.Model(&models.Configuration{}).Where("id = ?", configurationID).Delete(&models.Configuration{}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) GetConfigurationBySubtaskID(subtaskID string) (item []models.Configuration, err error) {
	err = b.db.Model(&models.Configuration{}).Where("subtask_id = ?", subtaskID).Find(&item).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}
