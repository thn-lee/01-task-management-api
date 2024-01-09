package tasks

import models "github.com/thn-lee/01-task-management-api/pkg/models/api"

func (u *taskUsecase) EditTask(taskID uint, requestTask models.TaskBody) (result models.Task, err error) {
	task := models.Task{
		Title:       requestTask.Title,
		Description: requestTask.Description,
	}

	return u.taskRepo.EditTask(taskID, task)
}
