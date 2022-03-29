package base

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/utils"

	"log"
)

func (b *Base) CreateTask(name string, description string) error {
	err := b.db.Model(&models.Task{}).Create(&models.Task{
		BaseModel:   models.BaseModel{ID: utils.GenerateID()},
		Name:        name,
		Token:       utils.RandKey(6),
		Description: description,
	}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) GetTasks() (tasks []models.Task, err error) {
	err = b.db.Model(&models.Task{}).Find(&tasks).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}

func (b *Base) UpdateTask(taskID string, name string, token string, description string) error {
	err := b.db.Model(&models.Task{}).Where("id = ?", taskID).Updates(&models.Task{
		Name:        name,
		Token:       token,
		Description: description,
	}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) DelTask(taskID string) (err error) {
	begin := b.db.Begin()
	defer func() {
		if err == nil {
			begin.Commit()
		} else {
			begin.Rollback()
		}
	}()

	err = begin.Model(&models.Task{}).Where("id = ?", taskID).Delete(&models.Task{}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	err = begin.Model(&models.Subtasks{}).Where("task_id = ?", taskID).Delete(&models.Subtasks{}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) CreateSubtasks(taskID string, name string, agentID string, branch string, instruction string, description string) error {
	b.db.Model(&models.Subtasks{}).Create(&models.Subtasks{
		BaseModel: models.BaseModel{ID: utils.GenerateID()},
	})
}
