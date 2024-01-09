// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
)

// TaskUsecase is an autogenerated mock type for the TaskUsecase type
type TaskUsecase struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: task
func (_m *TaskUsecase) CreateTask(task models.TaskBody) (models.Task, error) {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(models.TaskBody) (models.Task, error)); ok {
		return rf(task)
	}
	if rf, ok := ret.Get(0).(func(models.TaskBody) models.Task); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	if rf, ok := ret.Get(1).(func(models.TaskBody) error); ok {
		r1 = rf(task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTask provides a mock function with given fields: taskID
func (_m *TaskUsecase) DeleteTask(taskID uint) error {
	ret := _m.Called(taskID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(taskID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditTask provides a mock function with given fields: taskID, requestTask
func (_m *TaskUsecase) EditTask(taskID uint, requestTask models.TaskBody) (models.Task, error) {
	ret := _m.Called(taskID, requestTask)

	if len(ret) == 0 {
		panic("no return value specified for EditTask")
	}

	var r0 models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, models.TaskBody) (models.Task, error)); ok {
		return rf(taskID, requestTask)
	}
	if rf, ok := ret.Get(0).(func(uint, models.TaskBody) models.Task); ok {
		r0 = rf(taskID, requestTask)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	if rf, ok := ret.Get(1).(func(uint, models.TaskBody) error); ok {
		r1 = rf(taskID, requestTask)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditTaskStatus provides a mock function with given fields: taskID, requestTask
func (_m *TaskUsecase) EditTaskStatus(taskID uint, requestTask models.TaskBody) (models.Task, error) {
	ret := _m.Called(taskID, requestTask)

	if len(ret) == 0 {
		panic("no return value specified for EditTaskStatus")
	}

	var r0 models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, models.TaskBody) (models.Task, error)); ok {
		return rf(taskID, requestTask)
	}
	if rf, ok := ret.Get(0).(func(uint, models.TaskBody) models.Task); ok {
		r0 = rf(taskID, requestTask)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	if rf, ok := ret.Get(1).(func(uint, models.TaskBody) error); ok {
		r1 = rf(taskID, requestTask)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTask provides a mock function with given fields: taskID
func (_m *TaskUsecase) GetTask(taskID uint) (models.Task, error) {
	ret := _m.Called(taskID)

	if len(ret) == 0 {
		panic("no return value specified for GetTask")
	}

	var r0 models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (models.Task, error)); ok {
		return rf(taskID)
	}
	if rf, ok := ret.Get(0).(func(uint) models.Task); ok {
		r0 = rf(taskID)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTasks provides a mock function with given fields: criteria
func (_m *TaskUsecase) ListTasks(criteria models.Task) ([]models.Task, error) {
	ret := _m.Called(criteria)

	if len(ret) == 0 {
		panic("no return value specified for ListTasks")
	}

	var r0 []models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Task) ([]models.Task, error)); ok {
		return rf(criteria)
	}
	if rf, ok := ret.Get(0).(func(models.Task) []models.Task); ok {
		r0 = rf(criteria)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Task) error); ok {
		r1 = rf(criteria)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTaskUsecase creates a new instance of TaskUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskUsecase {
	mock := &TaskUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}