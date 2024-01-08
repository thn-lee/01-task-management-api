package tasks_test

// import (
// 	"errors"
// 	"testing"

// 	"github.com/brianvoe/gofakeit/v6"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/thn-lee/01-task-management-api/mocks"
// 	"github.com/thn-lee/01-task-management-api/pkg/models"
// 	"github.com/thn-lee/01-task-management-api/pkg/tasks"
// )

// func TestGetTaskUsecase(t *testing.T) {
// 	var mockTask models.Task
// 	gofakeit.Struct(&mockTask)

// 	mockTaskRepo := new(mocks.TaskRepository)
// 	mockTaskUsecase := new(mocks.TaskUsecase)

// 	t.Run("success", func(t *testing.T) {
// 		mockTaskRepo.On("GetTask", mock.AnythingOfType("string")).Return(mockTask, nil).Once()

// 		usecase := tasks.NewTaskUsecase(mockTaskRepo)

// 		result, err := usecase.GetTask(mockTask.ID)

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockTask, result)

// 		mockTaskUsecase.AssertExpectations(t)
// 	})
// 	t.Run("error-failed", func(t *testing.T) {
// 		mockTaskRepo.On("GetTask", mock.AnythingOfType("string")).Return(models.Task{}, errors.New("call error")).Once()

// 		usecase := tasks.NewTaskUsecase(mockTaskRepo)

// 		result, err := usecase.GetTask(mockTask.ID)

// 		assert.Error(t, err)
// 		assert.Equal(t, models.Task{}, result)

// 		mockTaskUsecase.AssertExpectations(t)
// 	})

// }

// func TestGetTasksUsecase(t *testing.T) {
// 	var mockTask models.Task
// 	gofakeit.Struct(&mockTask)

// 	mockTasks := []models.Task{mockTask}

// 	mockTaskRepo := new(mocks.TaskRepository)
// 	mockTaskUsecase := new(mocks.TaskUsecase)

// 	t.Run("success", func(t *testing.T) {
// 		mockTaskRepo.On("GetTasks", mock.Anything).Return(mockTasks, nil).Once()

// 		usecase := tasks.NewTaskUsecase(mockTaskRepo)

// 		result, err := usecase.GetTasks(mockTask)

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockTasks, result)

// 		mockTaskUsecase.AssertExpectations(t)
// 	})
// 	t.Run("error-failed", func(t *testing.T) {
// 		mockTaskRepo.On("GetTasks", mock.Anything).Return([]models.Task{}, errors.New("call error")).Once()

// 		usecase := tasks.NewTaskUsecase(mockTaskRepo)

// 		result, err := usecase.GetTasks(mockTask)

// 		assert.Error(t, err)
// 		assert.Empty(t, result)

// 		mockTaskUsecase.AssertExpectations(t)
// 	})

// }

// func TestCreateTaskUsecase(t *testing.T) {
// 	var mockTask models.Task
// 	gofakeit.Struct(&mockTask)

// 	mockTaskRepo := new(mocks.TaskRepository)
// 	mockTaskUsecase := new(mocks.TaskUsecase)

// 	t.Run("success", func(t *testing.T) {
// 		mockTaskRepo.On("CreateTask", mock.Anything).Return(nil).Once()

// 		usecase := tasks.NewTaskUsecase(mockTaskRepo)

// 		err := usecase.CreateTask(&mockTask)

// 		assert.NoError(t, err)

// 		mockTaskUsecase.AssertExpectations(t)
// 	})
// 	t.Run("error-failed", func(t *testing.T) {
// 		mockTaskRepo.On("CreateTask", mock.Anything).Return(errors.New("call error")).Once()

// 		usecase := tasks.NewTaskUsecase(mockTaskRepo)

// 		err := usecase.CreateTask(&mockTask)

// 		assert.Error(t, err)

// 		mockTaskUsecase.AssertExpectations(t)
// 	})

// }

// func TestEditTaskUsecase(t *testing.T) {
// 	var mockTask models.Task
// 	gofakeit.Struct(&mockTask)

// 	mockTaskRepo := new(mocks.TaskRepository)
// 	mockTaskUsecase := new(mocks.TaskUsecase)

// 	t.Run("success", func(t *testing.T) {
// 		mockTaskRepo.On("EditTask", mock.AnythingOfType("string"), mock.Anything).Return(nil).Once()

// 		usecase := tasks.NewTaskUsecase(mockTaskRepo)

// 		err := usecase.EditTask(mockTask.ID, mockTask)

// 		assert.NoError(t, err)

// 		mockTaskUsecase.AssertExpectations(t)
// 	})
// 	t.Run("error-failed", func(t *testing.T) {
// 		mockTaskRepo.On("EditTask", mock.AnythingOfType("string"), mock.Anything).Return(errors.New("call error")).Once()

// 		usecase := tasks.NewTaskUsecase(mockTaskRepo)

// 		err := usecase.EditTask(mockTask.ID, mockTask)

// 		assert.Error(t, err)

// 		mockTaskUsecase.AssertExpectations(t)
// 	})

// }

// func TestDeleteTaskUsecase(t *testing.T) {
// 	var mockTask models.Task
// 	gofakeit.Struct(&mockTask)

// 	mockTaskRepo := new(mocks.TaskRepository)
// 	mockTaskUsecase := new(mocks.TaskUsecase)

// 	t.Run("success", func(t *testing.T) {
// 		mockTaskRepo.On("DeleteTask", mock.AnythingOfType("string")).Return(nil).Once()

// 		usecase := tasks.NewTaskUsecase(mockTaskRepo)

// 		err := usecase.DeleteTask(mockTask.ID)

// 		assert.NoError(t, err)

// 		mockTaskUsecase.AssertExpectations(t)
// 	})
// 	t.Run("error-failed", func(t *testing.T) {
// 		mockTaskRepo.On("DeleteTask", mock.AnythingOfType("string")).Return(errors.New("call error")).Once()

// 		usecase := tasks.NewTaskUsecase(mockTaskRepo)

// 		err := usecase.DeleteTask(mockTask.ID)

// 		assert.Error(t, err)

// 		mockTaskUsecase.AssertExpectations(t)
// 	})

// }
