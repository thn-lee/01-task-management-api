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
	r.tasksMapListCounter.Lock()
	defer r.tasksMapListCounter.Unlock()
	if r.tasksMapListCounter.tasksMapList[int(taskID)] != (models.Task{}) {
		// update status
		if status != "" {
			updatedTask := r.tasksMapListCounter.tasksMapList[int(taskID)]
			updatedTask.Status = status
			updatedTask.UpdatedAt = time.Now()
			r.tasksMapListCounter.tasksMapList[int(taskID)] = updatedTask
		}
		return r.tasksMapListCounter.tasksMapList[int(taskID)], nil
	}

	// task not found
	log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
	err = fiber.NewError(http.StatusInternalServerError, "task not found.")
	return
}
