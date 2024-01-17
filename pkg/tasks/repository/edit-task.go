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
	// for taskIndex, task := range r.tasksMapList {
	r.tasksMapListCounter.Lock()
	defer r.tasksMapListCounter.Unlock()
	if r.tasksMapListCounter.tasksMapList[int(taskID)] != (models.Task{}) {
		// update title
		if incomingTask.Title != "" {
			updatedTask := r.tasksMapListCounter.tasksMapList[int(taskID)]
			updatedTask.Title = incomingTask.Title
			updatedTask.UpdatedAt = time.Now()
			r.tasksMapListCounter.tasksMapList[int(taskID)] = updatedTask
		}
		// update title description
		if incomingTask.Description != "" {
			updatedTask := r.tasksMapListCounter.tasksMapList[int(taskID)]
			updatedTask.Description = incomingTask.Description
			updatedTask.UpdatedAt = time.Now()
			r.tasksMapListCounter.tasksMapList[int(taskID)] = updatedTask
		}
		return r.tasksMapListCounter.tasksMapList[int(taskID)], nil
	}
	// }

	// task not found
	log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
	err = fiber.NewError(http.StatusInternalServerError, "task not found.")
	return
}
