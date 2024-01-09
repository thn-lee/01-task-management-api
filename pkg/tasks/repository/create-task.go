package tasks

import (
	"time"

	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
)

func (r *taskRepository) CreateTask(task models.Task) (result models.Task, err error) {
	task.ID = uint(len(r.taskList) + 1)
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	r.taskList = append(r.taskList, task)

	result = task
	return
}
