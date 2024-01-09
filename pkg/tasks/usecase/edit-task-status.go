package tasks

import (
	"errors"
	"log"

	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
)

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
