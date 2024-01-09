package tasks

import models "github.com/thn-lee/01-task-management-api/pkg/models/api"

func (r *taskRepository) ListTasks(criteria models.Task) (tasks []models.Task, err error) {
	tasks = r.taskList
	return
}
