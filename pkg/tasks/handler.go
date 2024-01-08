package tasks

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/thn-lee/01-task-management-api/pkg/domain"
	respModels "github.com/thn-lee/01-task-management-api/pkg/models"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
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

func (h *TaskHandler) GetTask() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := respModels.ResponseForm{}

		taskID, _ := c.ParamsInt("id")

		if taskID != 0 {
			task, err := h.TaskUsecase.GetTask(uint(taskID))
			if err != nil {
				responseForm.Success = false
				responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
					Code:    fiber.StatusBadRequest,
					Title:   http.StatusText(http.StatusBadRequest),
					Message: "Invalid Request.",
				})
				return c.Status(fiber.StatusBadRequest).JSON(responseForm)
			}

			responseForm.Result = map[string]interface{}{
				"task": task,
			}
		} else {
			var criteria models.Task
			if err := c.QueryParser(&criteria); err != nil {
				responseForm.Success = false
				responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
					Code:    fiber.StatusBadRequest,
					Title:   http.StatusText(http.StatusBadRequest),
					Message: fmt.Sprintf("Invalid Request, with err: %+v", err.Error()),
				})
				return c.Status(fiber.StatusBadRequest).JSON(responseForm)
			}

			tasks, err := h.TaskUsecase.GetTasks(criteria)
			if err != nil {
				return nil
			}

			responseForm.Result = map[string]interface{}{
				"tasks": tasks,
			}
		}

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}

func (h *TaskHandler) CreateTask() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := respModels.ResponseForm{}

		var requestBody models.TaskBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: fmt.Sprintf("Invalid Request, with err: %+v", err.Error()),
			})
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}

		task, err := h.TaskUsecase.CreateTask(&requestBody)
		if err != nil {
			log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		responseForm.Result = map[string]interface{}{
			"task": task,
		}

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}

func (h *TaskHandler) EditTask() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := respModels.ResponseForm{}

		taskID, err := c.ParamsInt("id")
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: fmt.Sprintf("Invalid Request, with err: %+v", err.Error()),
			})
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}

		var requestBody models.TaskBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: fmt.Sprintf("Invalid Request, with err: %+v", err.Error()),
			})
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}

		task, err := h.TaskUsecase.EditTask(uint(taskID), requestBody)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		responseForm.Result = map[string]interface{}{
			"task": task,
		}

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}

func (h *TaskHandler) EditTaskStatus() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := respModels.ResponseForm{}

		taskID, err := c.ParamsInt("id")
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: fmt.Sprintf("Invalid Request, with err: %+v", err.Error()),
			})
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}

		var requestBody models.TaskBody
		if err := c.BodyParser(&requestBody); err != nil {
			log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: fmt.Sprintf("Invalid Request, with err: %+v", err.Error()),
			})
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}

		task, err := h.TaskUsecase.EditTaskStatus(uint(taskID), requestBody)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		responseForm.Result = map[string]interface{}{
			"task": task,
		}

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}

func (h *TaskHandler) DeleteTask() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := respModels.ResponseForm{}

		taskID, err := c.ParamsInt("id")
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: fmt.Sprintf("Invalid Request, with err: %+v", err.Error()),
			})
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}

		if err := h.TaskUsecase.DeleteTask(uint(taskID)); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}
