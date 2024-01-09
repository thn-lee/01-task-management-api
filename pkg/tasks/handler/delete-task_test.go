package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thn-lee/01-task-management-api/pkg/domain/mocks"
	respModels "github.com/thn-lee/01-task-management-api/pkg/models"
	tasksUsecase "github.com/thn-lee/01-task-management-api/pkg/tasks/usecase"
)

func TestDeleteTaskHandler(t *testing.T) {
	// gofakeit.Struct(&mockTask)
	taskID := uint(1)

	mockResponse := respModels.ResponseForm{
		Success: true,
	}

	mockRepo := new(mocks.TaskRepository)
	mockRepo.On("DeleteTask", taskID).Return(nil)

	app := fiber.New()
	req, err := http.NewRequest("DELETE", "/api/v1/task/"+fmt.Sprintf("%+v", taskID), nil)
	assert.NoError(t, err)

	mockUsecase := tasksUsecase.NewTaskUsecase(mockRepo)
	NewTaskHandler(app.Group("/api/v1/task"), mockUsecase)

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockJson, err := json.Marshal(mockResponse)
	assert.NoError(t, err)

	var bodyBytes bytes.Buffer
	io.Copy(&bodyBytes, resp.Body)

	assert.JSONEq(t, string(mockJson), bodyBytes.String())

	mockRepo.AssertExpectations(t)
}

func TestBadRequestTaskHandler(t *testing.T) {
	// gofakeit.Struct(&mockTask)

	mockResponse := respModels.ResponseForm{
		Success: false,
		Errors: []respModels.ResponseError{
			{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(fiber.StatusBadRequest),
				Message: "Invalid Request.",
			},
		},
	}

	mockRepo := new(mocks.TaskRepository)

	app := fiber.New()
	req, err := http.NewRequest("DELETE", "/api/v1/task/string", nil)
	assert.NoError(t, err)

	mockUsecase := tasksUsecase.NewTaskUsecase(mockRepo)
	NewTaskHandler(app.Group("/api/v1/task"), mockUsecase)

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	mockJson, err := json.Marshal(mockResponse)
	assert.NoError(t, err)

	var bodyBytes bytes.Buffer
	io.Copy(&bodyBytes, resp.Body)

	assert.JSONEq(t, string(mockJson), bodyBytes.String())

	mockRepo.AssertExpectations(t)
}

func TestTaskNotFoundHandler(t *testing.T) {
	taskID := uint(99)
	// gofakeit.Struct(&mockTask)

	mockResponse := respModels.ResponseForm{
		Success: false,
		Errors: []respModels.ResponseError{
			{
				Code:    fiber.StatusInternalServerError,
				Title:   http.StatusText(fiber.StatusInternalServerError),
				Message: "task not found.",
			},
		},
	}

	mockRepo := new(mocks.TaskRepository)
	mockRepo.On("DeleteTask", taskID).Return(fiber.NewError(http.StatusInternalServerError, "task not found."))

	app := fiber.New()
	req, err := http.NewRequest("DELETE", "/api/v1/task/99", nil)
	assert.NoError(t, err)

	mockUsecase := tasksUsecase.NewTaskUsecase(mockRepo)
	NewTaskHandler(app.Group("/api/v1/task"), mockUsecase)

	resp, err := app.Test(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	mockJson, err := json.Marshal(mockResponse)
	assert.NoError(t, err)

	var bodyBytes bytes.Buffer
	io.Copy(&bodyBytes, resp.Body)

	assert.JSONEq(t, string(mockJson), bodyBytes.String())

	mockRepo.AssertExpectations(t)
}
