package tasks

import (
	"time"

	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
)

func (r *taskRepository) CreateTask(task models.Task) (result models.Task, err error) {
	r.tasksMapListCounter.Lock()
	defer r.tasksMapListCounter.Unlock()
	latestID := len(r.tasksMapListCounter.tasksMapList)
	r.tasksMapListCounter.tasksMapList[latestID+1] = models.Task{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		ID:          uint(latestID + 1),
	}

	result = r.tasksMapListCounter.tasksMapList[latestID+1]
	return
}
