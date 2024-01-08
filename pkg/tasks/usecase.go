package tasks

import (
	"errors"
	"log"

	"github.com/thn-lee/01-task-management-api/pkg/domain"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
	// "github.com/thn-lee/01-task-management-api/pkg/models"
)

type taskUsecase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUsecase(taskRepo domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{
		taskRepo: taskRepo,
	}
}

const (
	StatusToDo       = iota // 0
	StatusInProgress        // 1
	StatusDone              // 2
)

func (u *taskUsecase) GetTask(taskID uint) (task models.Task, err error) {
	return u.taskRepo.GetTask(taskID)
}

func (u *taskUsecase) GetTasks(criteria models.Task) (tasks []models.Task, err error) {
	return u.taskRepo.GetTasks(criteria)
}

func (u *taskUsecase) CreateTask(requestTask *models.TaskBody) (result models.Task, err error) {
	task := models.Task{
		Title:       requestTask.Title,
		Description: requestTask.Description,
	}

	switch requestTask.Status {
	case StatusToDo:
		task.Status = "To Do"
		return u.taskRepo.CreateTask(&task)
	case StatusInProgress:
		task.Status = "In Progress"
		return u.taskRepo.CreateTask(&task)
	case StatusDone:
		task.Status = "Done"
		return u.taskRepo.CreateTask(&task)
	default:
		log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
		return models.Task{}, errors.New("invalid status")
	}
}

func (u *taskUsecase) EditTask(taskID uint, requestTask models.TaskBody) (result models.Task, err error) {
	task := models.Task{
		Title:       requestTask.Title,
		Description: requestTask.Description,
	}

	return u.taskRepo.EditTask(taskID, task)
}

func (u *taskUsecase) EditTaskStatus(taskID uint, requestTask models.TaskBody) (result models.Task, err error) {
	switch requestTask.Status {
	case StatusToDo:
		status := "To Do"
		return u.taskRepo.EditTaskStatus(taskID, status)
	case StatusInProgress:
		status := "In Progress"
		return u.taskRepo.EditTaskStatus(taskID, status)
	case StatusDone:
		status := "Done"
		return u.taskRepo.EditTaskStatus(taskID, status)
	default:
		log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
		return models.Task{}, errors.New("invalid status")
	}
}

func (u *taskUsecase) DeleteTask(taskID uint) (err error) {
	return u.taskRepo.DeleteTask(taskID)
}
