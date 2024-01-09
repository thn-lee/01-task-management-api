package tasks

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
)

func (r *taskRepository) EditTask(taskID uint, incomingTask models.Task) (result models.Task, err error) {
	for taskIndex, task := range r.taskList {
		if task.ID == taskID {
			// update title
			if incomingTask.Title != "" {
				r.taskList[taskIndex].Title = incomingTask.Title
			}
			// update title description
			if incomingTask.Description != "" {
				r.taskList[taskIndex].Description = incomingTask.Description
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
