package tasks

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
)

func (r *taskRepository) DeleteTask(taskID uint) (err error) {
	for taskIndex, task := range r.taskList {
		if task.ID == taskID {
			r.taskList = append(r.taskList[:taskIndex], r.taskList[taskIndex+1:]...)
			return nil
		}
	}

	// task not found
	log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
	err = fiber.NewError(http.StatusInternalServerError, "task not found.")
	return
}
