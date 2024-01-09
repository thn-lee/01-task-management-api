package tasks

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
)

func (r *taskRepository) EditTaskStatus(taskID uint, status string) (result models.Task, err error) {
	for taskIndex, task := range r.taskList {
		if task.ID == taskID {
			// update status
			if status != "" {
				r.taskList[taskIndex].Status = status
			}
			r.taskList[taskIndex].UpdatedAt = time.Now()
			return r.taskList[taskIndex], nil
		}
	}

	// task not found
	log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
	err = fiber.NewError(http.StatusInternalServerError, "task not found.")
	return
}
