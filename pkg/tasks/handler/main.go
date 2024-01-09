package tasks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thn-lee/01-task-management-api/pkg/domain"
)

type TaskHandler struct {
	TaskUsecase domain.TaskUsecase
}

// NewPostHandler will initialize the post resource endpoint
func NewTaskHandler(router fiber.Router, taskUsecase domain.TaskUsecase) {

	handler := &TaskHandler{
		TaskUsecase: taskUsecase,
	}

	router.Get("/:id?", handler.GetTask())
	router.Post("/", handler.CreateTask())
	router.Patch("/:id", handler.EditTask())
	router.Patch("/:id/status", handler.EditTaskStatus())
	router.Delete("/:id", handler.DeleteTask())
}
