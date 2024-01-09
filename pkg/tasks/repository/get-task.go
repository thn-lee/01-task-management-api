package tasks

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
)

func (r *taskRepository) GetTask(taskID uint) (task models.Task, err error) {
	for _, task := range r.taskList {
		if task.ID == taskID {
			return task, nil
		}
	}

	log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
	err = fiber.NewError(http.StatusInternalServerError, "task not found.")
	return
}
