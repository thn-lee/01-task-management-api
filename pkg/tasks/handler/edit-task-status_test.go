package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thn-lee/01-task-management-api/pkg/domain/mocks"
	respModels "github.com/thn-lee/01-task-management-api/pkg/models"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	tasksUsecase "github.com/thn-lee/01-task-management-api/pkg/tasks/usecase"
)

func TestEditTaskStatusHandler(t *testing.T) {
	mockTask := models.Task{ID: 1, Status: "In Progress"}
	// gofakeit.Struct(&mockTask)
	payload := strings.NewReader(`{
		"status": 2
	}`)

	mockResponse := respModels.ResponseForm{
		Success: true,
		Result: map[string]interface{}{
			"task": mockTask,
		},
	}

	mockRepo := new(mocks.TaskRepository)
	mockRepo.On("EditTaskStatus", mockTask.ID, "Done").Return(mockTask, nil)

	app := fiber.New()
	req, err := http.NewRequest("PATCH", "/api/v1/task/"+fmt.Sprintf("%+v/status", mockTask.ID), payload)
	req.Header.Set("Content-Type", "application/json")

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

func TestErrBadRequestTaskIDStatusHandler(t *testing.T) {
	// mockTask := models.Task{ID: 1, Title: "Updated Title", Description: "Updated Description"}
	// gofakeit.Struct(&mockTask)
	payload := strings.NewReader(`{
		"status": 2
	}`)

	mockResponse := respModels.ResponseForm{
		Success: false,
		Errors: []respModels.ResponseError{
			{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(fiber.StatusBadRequest),
				Message: "Invalid Request. taskID must be an integer.",
			},
		},
	}

	mockRepo := new(mocks.TaskRepository)

	app := fiber.New()
	req, err := http.NewRequest("PATCH", "/api/v1/task/string/status", payload)
	req.Header.Set("Content-Type", "application/json")

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

func TestErrBadRequestBodyParserStatusHandler(t *testing.T) {
	mockTask := models.Task{ID: 1, Title: "Updated Title", Description: "Updated Description"}
	// gofakeit.Struct(&mockTask)
	payload := strings.NewReader(`{
		"status": 2
	}`)

	mockResponse := respModels.ResponseForm{
		Success: false,
		Errors: []respModels.ResponseError{
			{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(fiber.StatusBadRequest),
				Message: "Invalid Request. Cannot parsing request body.",
			},
		},
	}

	mockRepo := new(mocks.TaskRepository)

	app := fiber.New()
	req, err := http.NewRequest("PATCH", "/api/v1/task/"+fmt.Sprintf("%+v/status", mockTask.ID), payload)
	// req.Header.Set("Content-Type", "application/json")

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

func TestNotFoundTaskStatusHandler(t *testing.T) {
	mockTask := models.Task{ID: 99, Title: "Updated Title", Description: "Updated Description"}
	// gofakeit.Struct(&mockTask)
	payload := strings.NewReader(`{
		"status": 2
	}`)

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
	mockRepo.On("EditTaskStatus", mockTask.ID, "Done").Return(models.Task{}, fiber.NewError(http.StatusInternalServerError, "task not found."))

	app := fiber.New()
	req, err := http.NewRequest("PATCH", "/api/v1/task/"+fmt.Sprintf("%+v/status", mockTask.ID), payload)
	req.Header.Set("Content-Type", "application/json")

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
