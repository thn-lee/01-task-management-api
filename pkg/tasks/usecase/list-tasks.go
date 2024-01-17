package tasks

import (
	"errors"
	"log"

	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
)

func (u *taskUsecase) ListTasks(criteria models.TaskBody) (tasks map[int]models.Task, err error) {

	switch criteria.Status {
	case StatusToDo:
		status := "To Do"
		task := models.Task{
			Title:  criteria.Title,
			Status: status,
		}
		return u.taskRepo.ListTasks(task)
	case StatusInProgress:
		status := "In Progress"
		task := models.Task{
			Title:  criteria.Title,
			Status: status,
		}
		return u.taskRepo.ListTasks(task)
	case StatusDone:
		status := "Done"
		task := models.Task{
			Title:  criteria.Title,
			Status: status,
		}
		return u.taskRepo.ListTasks(task)
	default:
		log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
		return tasks, errors.New("invalid status")
	}
}
