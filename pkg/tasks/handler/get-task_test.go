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
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"

	tasksUsecase "github.com/thn-lee/01-task-management-api/pkg/tasks/usecase"
)

func TestGetTaskHandler(t *testing.T) {
	mockTask := models.Task{ID: 1, Title: "Test Title", Description: "Test Description", Status: "Test Status"}
	// gofakeit.Struct(&mockTask)

	mockResponse := respModels.ResponseForm{
		Success: true,
		Result: map[string]interface{}{
			"task": mockTask,
		},
	}

	mockRepo := new(mocks.TaskRepository)

	// GetTask(taskID uint) (task models.Task, err error)
	mockRepo.On("GetTask", uint(mockTask.ID)).Return(mockTask, nil)

	app := fiber.New()
	req, err := http.NewRequest("GET", "/api/v1/task/"+fmt.Sprintf("%+v", mockTask.ID), nil)
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

func TestListTaskHandler(t *testing.T) {
	mockTask := []models.Task{{ID: 1, Title: "Test Title", Description: "Test Description", Status: "Test Status"}}
	var criteria models.Task
	// gofakeit.Struct(&mockTask)

	mockResponse := respModels.ResponseForm{
		Success: true,
		Result: map[string]interface{}{
			"tasks": mockTask,
		},
	}

	mockRepo := new(mocks.TaskRepository)

	// GetTask(taskID uint) (task models.Task, err error)
	mockRepo.On("ListTasks", criteria).Return(mockTask, nil)

	app := fiber.New()
	req, err := http.NewRequest("GET", "/api/v1/task/", nil)
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

func TestGetTaskBadRequestHandler(t *testing.T) {
	mockResponse := respModels.ResponseForm{
		Success: false,
		Errors: []respModels.ResponseError{
			{
				Code:    fiber.StatusBadRequest,
				Title:   http.StatusText(fiber.StatusBadRequest),
				Message: "Invalid Request. err: task not found",
			},
		},
	}

	mockRepo := new(mocks.TaskRepository)
	mockRepo.On("GetTask", uint(100)).Return(models.Task{}, fiber.NewError(http.StatusInternalServerError, "task not found"))

	app := fiber.New()
	req, err := http.NewRequest("GET", "/api/v1/task/100", nil)
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

// func TestListBadRequestHandler(t *testing.T) {
// 	// TODO: for future improvement to make this api searchable or filterable

// 	mockResponse := respModels.ResponseForm{
// 		Success: false,
// 		Errors: []respModels.ResponseError{
// 			{
// 				Code:    fiber.StatusBadRequest,
// 				Title:   http.StatusText(fiber.StatusBadRequest),
// 				Message: "Invalid Request. err: task not found",
// 			},
// 		},
// 	}

// 	mockRepo := new(mocks.TaskRepository)
// 	mockRepo.On("GetTask", uint(100)).Return(models.Task{}, fiber.NewError(http.StatusInternalServerError, "task not found"))

// 	app := fiber.New()
// 	req, err := http.NewRequest("GET", "/api/v1/task", nil)
// 	assert.NoError(t, err)

// 	mockUsecase := tasksUsecase.NewTaskUsecase(mockRepo)
// 	NewTaskHandler(app.Group("/api/v1/task"), mockUsecase)

// 	resp, err := app.Test(req)
// 	require.NoError(t, err)

// 	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

// 	mockJson, err := json.Marshal(mockResponse)
// 	assert.NoError(t, err)

// 	var bodyBytes bytes.Buffer
// 	io.Copy(&bodyBytes, resp.Body)

// 	assert.JSONEq(t, string(mockJson), bodyBytes.String())

// 	mockRepo.AssertExpectations(t)
// }
