package task_pool

import (
	"github.com/dollarkillerx/fireworks/internal/pkg/models"

	"sync"
)

type TaskPool struct {
	task map[string][]models.Subtasks
	mu   sync.Mutex
}

func New() *TaskPool {
	return &TaskPool{
		task: map[string][]models.Subtasks{},
	}
}

func (t *TaskPool) AddTask(pool string, task models.Subtasks) {
	t.mu.Lock()
	defer t.mu.Unlock()

	subtasks, ex := t.task[pool]
	if !ex {
		subtasks = make([]models.Subtasks, 0)
	}

	subtasks = append(subtasks, task)

	t.task[pool] = subtasks
}

func (t *TaskPool) GetTasks(pool string) []models.Subtasks {
	t.mu.Lock()
	defer t.mu.Unlock()

	subtasks := t.task[pool]

	t.task[pool] = make([]models.Subtasks, 0)

	return subtasks
}
