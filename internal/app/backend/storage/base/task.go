package base

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
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
	err = b.db.Model(&models.Task{}).Order("updated_at desc").Find(&tasks).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}

func (b *Base) DisabledTask(taskID string) error {
	var task models.Task
	err := b.db.Model(&models.Task{}).Where("id = ?", taskID).First(&task).Error
	if err != nil {
		log.Println(err)
		return err
	}

	err = b.db.Model(&models.Task{}).Where("id = ?", taskID).
		Update("disabled", !task.Disabled).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
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

	err = begin.Model(&models.Task{}).Where("id = ?", taskID).Unscoped().Delete(&models.Task{}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	err = begin.Model(&models.Subtasks{}).Where("task_id = ?", taskID).Unscoped().Delete(&models.Subtasks{}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) GetTaskByToken(token string) (*models.Task, error) {
	var task models.Task
	err := b.db.Model(&models.Task{}).Where("token = ?", token).First(&task).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &task, nil
}

func (b *Base) CreateSubtasks(taskID string, name string, agentID string, branch string, action enum.TaskAction, instruction string, description string) error {
	err := b.db.Model(&models.Subtasks{}).Create(&models.Subtasks{
		BaseModel:   models.BaseModel{ID: utils.GenerateID()},
		Name:        name,
		TaskID:      taskID,
		AgentID:     agentID,
		Branch:      branch,
		Action:      action,
		Instruction: instruction,
		Description: description,
	}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) DisabledSubtasks(subtasksID string) error {
	var sub models.Subtasks
	err := b.db.Model(&models.Subtasks{}).Where("id = ?", subtasksID).First(&sub).Error
	if err != nil {
		log.Println(err)
		return err
	}

	err = b.db.Model(&models.Subtasks{}).Where("id = ?", subtasksID).
		Update("disabled", !sub.Disabled).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) DelSubtasks(subtasksID string) error {
	err := b.db.Model(&models.Subtasks{}).Where("id = ?", subtasksID).Unscoped().Delete(&models.Subtasks{}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) GetSubtasks(taskID string) (subs []models.Subtasks, err error) {
	err = b.db.Model(&models.Subtasks{}).Where("task_id = ?", taskID).Order("updated_at desc").Find(&subs).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}

func (b *Base) UpdateSubtasks(subtaskID string, instruction string, description string) error {
	err := b.db.Model(&models.Subtasks{}).Where("id = ?", subtaskID).Updates(&models.Subtasks{
		Instruction: instruction,
		Description: description,
	}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *Base) GetSubtasksBySubtasksID(subtaskID string) (*models.Subtasks, error) {
	var sub models.Subtasks
	err := b.db.Model(&models.Subtasks{}).Where("id = ?", subtaskID).First(&sub).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &sub, nil
}

func (b *Base) CreateTaskLog(subtasksID string, taskType enum.TaskType) (id string, err error) {
	id = utils.GenerateID()
	err = b.db.Model(&models.TaskLog{}).Create(&models.TaskLog{
		BaseModel:  models.BaseModel{ID: id},
		SubtasksID: subtasksID,
		TaskStage:  enum.TaskStageBuild,
		TaskStatus: enum.TaskStatusWait,
		TaskType:   taskType,
	}).Error
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (b *Base) UpdateTaskLog(logID string, taskStatus enum.TaskStatus, taskStage enum.TaskStage, logText string) error {
	err := b.db.Model(&models.TaskLog{}).
		Where("id = ?", logID).Updates(&models.TaskLog{
		TaskStage:  taskStage,
		TaskStatus: taskStatus,
		Log:        logText,
	}).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return nil
}

func (b *Base) GetTaskLog(subtasksID string) (logs []models.TaskLog, err error) {
	err = b.db.Model(&models.TaskLog{}).
		Where("subtasks_id = ?", subtasksID).Order("updated_at desc").Limit(13).Find(&logs).Error
	if err != nil {
		log.Println(err)
		return
	}

	return
}
