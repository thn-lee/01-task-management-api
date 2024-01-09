package tasks

import models "github.com/thn-lee/01-task-management-api/pkg/models/api"

func (u *taskUsecase) GetTask(taskID uint) (task models.Task, err error) {
	return u.taskRepo.GetTask(taskID)
}
