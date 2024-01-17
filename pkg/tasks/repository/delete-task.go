package tasks

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
)

func (r *taskRepository) DeleteTask(taskID uint) (err error) {
	r.tasksMapListCounter.Lock()
	defer r.tasksMapListCounter.Unlock()
	if r.tasksMapListCounter.tasksMapList[int(taskID)] != (models.Task{}) {
		delete(r.tasksMapListCounter.tasksMapList, int(taskID))
		return nil
	}

	// task not found
	log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
	err = fiber.NewError(http.StatusInternalServerError, "task not found.")
	return
}
