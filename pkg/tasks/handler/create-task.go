package tasks

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	respModels "github.com/thn-lee/01-task-management-api/pkg/models"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"
)

func (h *TaskHandler) CreateTask() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := respModels.ResponseForm{}

		var requestBody models.TaskBody
		if err := c.BodyParser(&requestBody); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: fmt.Sprintf("Invalid Request, with err: %+v", err.Error()),
			})
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}

		task, err := h.TaskUsecase.CreateTask(requestBody)
		if err != nil {
			log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, respModels.ResponseError{
				Code:    fiber.StatusInternalServerError,
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
