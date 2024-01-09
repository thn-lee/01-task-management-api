package tasks

import models "github.com/thn-lee/01-task-management-api/pkg/models/api"

func (u *taskUsecase) ListTasks(criteria models.Task) (tasks []models.Task, err error) {
	return u.taskRepo.ListTasks(criteria)
}
