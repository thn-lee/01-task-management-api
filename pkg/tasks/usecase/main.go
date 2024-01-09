package tasks

import (
	"github.com/thn-lee/01-task-management-api/pkg/domain"
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
