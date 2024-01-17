package tasks

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	respModels "github.com/thn-lee/01-task-management-api/pkg/models"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
)

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
					Message: fmt.Sprintf("Invalid Request. err: %+v", err.Error()),
				})
				return c.Status(fiber.StatusBadRequest).JSON(responseForm)
			}

			responseForm.Result = map[string]interface{}{
				"task": task,
			}
		} else {
			// {{base_urk}}/api/v1/tasks/1?criteria="some_text", order="", page limit ...
			// TODO: for future improvement to make this api searchable or filterable
			var criteria models.TaskBody
			if err := c.QueryParser(&criteria); err != nil {
				responseForm.Success = false
				responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
					Code:    fiber.StatusBadRequest,
					Title:   http.StatusText(http.StatusBadRequest),
					Message: fmt.Sprintf("Invalid Request, Cannot parse query with err: %+v", err.Error()),
				})
				return c.Status(fiber.StatusBadRequest).JSON(responseForm)
			}

			tasks, err := h.TaskUsecase.ListTasks(criteria)
			if err != nil {
				responseForm.Success = false
				responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
					Code:    fiber.StatusInternalServerError,
					Title:   http.StatusText(http.StatusInternalServerError),
					Message: err.Error(),
				})
				return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
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
