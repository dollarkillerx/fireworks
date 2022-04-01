package base

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/enum"
	"github.com/dollarkillerx/fireworks/internal/pkg/models"
	"github.com/dollarkillerx/fireworks/internal/utils"

	"log"
	"time"
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

func (b *Base) GetTaskByID(taskID string) (*models.Task, error) {
	var task models.Task
	err := b.db.Model(&models.Task{}).Where("id = ?", taskID).First(&task).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &task, nil
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

	if sub.AgentID != "" {
		var agent models.Agent
		err := b.db.Model(&models.Agent{}).Where("id = ?", sub.AgentID).First(&agent).Error
		if err != nil {
			log.Println(err)
			return nil, err
		}

		sub.AgentName = agent.AgentName
	}

	return &sub, nil
}

func (b *Base) CreateTaskLog(subtasksID string, gitSSH string, taskType enum.TaskType) (id string, err error) {
	id = utils.GenerateID()

	subtasks, err := b.GetSubtasksBySubtasksID(subtasksID)
	if err != nil {
		return "", err
	}

	task, err := b.GetTaskByID(subtasks.TaskID)
	if err != nil {
		return "", err
	}

	if gitSSH != "" {
		err = b.db.Model(&models.Subtasks{}).Where("id = ?", subtasksID).Update("git_addr", gitSSH).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	err = b.db.Model(&models.TaskLog{}).Create(&models.TaskLog{
		BaseModel:   models.BaseModel{ID: id},
		TaskID:      subtasks.TaskID,
		TaskName:    task.Name,
		SubtasksID:  subtasksID,
		SubtaskName: subtasks.Name,
		TaskStage:   enum.TaskStageBuild,
		TaskStatus:  enum.TaskStatusWait,
		TaskType:    taskType,
		AgentID:     subtasks.AgentName,
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

func (b *Base) GetTaskLog(taskID string, subtasksID string) (logs []models.TaskLog, err error) {
	sql := b.db.Model(&models.TaskLog{})
	if taskID != "" {
		sql = sql.Where("task_id = ?", taskID)
	}
	if subtasksID != "" {
		sql = sql.Where("subtasks_id = ?", subtasksID)
	}
	err = sql.Order("updated_at desc").Limit(20).Find(&logs).Error
	if err != nil {
		log.Println(err)
		return
	}

	for i := range logs {
		logs[i].DisplayTime = logs[i].UpdatedAt.Format("2006-01-02 15:04:05")
	}

	return
}

// updateTaskOldLogs 内部调用 设置过期服务
func (b *Base) updateTaskOldLogs() {
	var logs []models.TaskLog
	err := b.db.Model(&models.TaskLog{}).Where("task_status in (?)", []interface{}{enum.TaskStatusWait, enum.TaskStatusRunning}).Find(&logs).Error
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range logs {
		if v.CreatedAt.Add(time.Minute*60).Unix() < time.Now().Unix() {
			err := b.db.Model(&models.TaskLog{}).Where("id = ?", v.ID).Updates(&models.TaskLog{
				TaskStatus: enum.TaskStatusFailed,
				Log:        "time out",
			}).Error
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// RecieveTaskByLog 通过日志获取任务
func (b *Base) RecieveTaskByLog(agentName string) (subs []models.Subtasks, err error) {
	var taskLogs []models.TaskLog
	err = b.db.Model(&models.TaskLog{}).
		Where("agent_id = ?", agentName).
		Where("task_status = ?", enum.TaskStatusWait).Find(&taskLogs).Error
	if err != nil {
		return nil, err
	}

	if len(taskLogs) == 0 {
		return
	}

	var subID []string
	for _, v := range taskLogs {
		subID = append(subID, v.SubtasksID)
	}

	err = b.db.Model(&models.Subtasks{}).Where("id in (?)", subID).Find(&subs).Error
	if err != nil {
		return nil, err
	}

	if len(subs) == 0 {
		return
	}

	for i, v := range subs {
		for _, vv := range taskLogs {
			if v.ID == vv.SubtasksID {
				subs[i].LogID = vv.ID
				subs[i].TaskType = vv.TaskType
			}
		}
	}

	return
}
