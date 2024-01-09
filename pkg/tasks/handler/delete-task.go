package tasks

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	respModels "github.com/thn-lee/01-task-management-api/pkg/models"
)

func (h *TaskHandler) DeleteTask() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := respModels.ResponseForm{}

		taskID, err := c.ParamsInt("id")
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: fmt.Sprintf("Invalid Request."),
			})
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}

		if err := h.TaskUsecase.DeleteTask(uint(taskID)); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusInternalServerError,
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
