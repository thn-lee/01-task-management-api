package tasks

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
)

func (u *taskUsecase) CreateTask(requestTask models.TaskBody) (result models.Task, err error) {
	task := models.Task{
		Title:       requestTask.Title,
		Description: requestTask.Description,
	}

	switch requestTask.Status {
	case StatusToDo:
		task.Status = "To Do"
		return u.taskRepo.CreateTask(task)
	case StatusInProgress:
		task.Status = "In Progress"
		return u.taskRepo.CreateTask(task)
	case StatusDone:
		task.Status = "Done"
		return u.taskRepo.CreateTask(task)
	default:
		log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
		return models.Task{}, fiber.NewError(http.StatusInternalServerError, "Invalid status.")
	}
}
