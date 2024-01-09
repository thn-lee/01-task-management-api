package domain

import models "github.com/thn-lee/01-task-management-api/pkg/models/api"

//go:generate mockery --name TaskUsecase
type TaskUsecase interface {
	GetTask(taskID uint) (task models.Task, err error)
	ListTasks(criteria models.Task) (tasks []models.Task, err error)
	CreateTask(task models.TaskBody) (result models.Task, err error)
	EditTask(taskID uint, requestTask models.TaskBody) (result models.Task, err error)
	EditTaskStatus(taskID uint, requestTask models.TaskBody) (result models.Task, err error)
	DeleteTask(taskID uint) (err error)
}

//go:generate mockery --name TaskRepository
type TaskRepository interface {
	GetTask(taskID uint) (task models.Task, err error)
	ListTasks(criteria models.Task) (tasks []models.Task, err error)
	CreateTask(task models.Task) (result models.Task, err error)
	EditTask(taskID uint, requestTask models.Task) (result models.Task, err error)
	EditTaskStatus(taskID uint, status string) (result models.Task, err error)

	DeleteTask(taskID uint) (err error)
}
