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

func TestCreateTaskHandler(t *testing.T) {
	mockTask := models.Task{Title: "Test Title", Description: "Test Description", Status: "In Progress"}
	// gofakeit.Struct(&mockTask)
	payload := strings.NewReader(`{
		"title": "Test Title",
		"description": "Test Description",
		"status": 1
	}`)

	mockResponse := respModels.ResponseForm{
		Success: true,
		Result: map[string]interface{}{
			"task": mockTask,
		},
	}

	mockRepo := new(mocks.TaskRepository)
	mockRepo.On("CreateTask", mockTask).Return(mockTask, nil)

	app := fiber.New()
	req, err := http.NewRequest("POST", "/api/v1/task", payload)
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

func TestErrBodyParserHandler(t *testing.T) {
	// gofakeit.Struct(&mockTask)
	payload := strings.NewReader(`{
		"title": "Test Title",
		"description": "Test Description",
		"status": 1
	}`)

	mockResponse := respModels.ResponseForm{
		Success: false,
		Errors: []respModels.ResponseError{
			{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(fiber.StatusBadRequest),
				Message: fmt.Sprintf("Invalid Request, with err: Unprocessable Entity"),
			},
		},
	}

	mockRepo := new(mocks.TaskRepository)

	app := fiber.New()
	req, err := http.NewRequest("POST", "/api/v1/task", payload)
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

func TestErrCreateTaskInvalidStatusHandler(t *testing.T) {
	// gofakeit.Struct(&mockTask)
	payload := strings.NewReader(`{
		"title": "Test Title",
		"description": "Test Description",
		"status": 99
	}`)

	mockResponse := respModels.ResponseForm{
		Success: false,
		Errors: []respModels.ResponseError{
			{
				Code:    fiber.StatusInternalServerError,
				Title:   http.StatusText(fiber.StatusInternalServerError),
				Message: fmt.Sprintf("Invalid status."),
			},
		},
	}

	mockRepo := new(mocks.TaskRepository)

	app := fiber.New()
	req, err := http.NewRequest("POST", "/api/v1/task", payload)
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
