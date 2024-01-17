package tasks

import (
	"sync"

	"github.com/thn-lee/01-task-management-api/pkg/domain"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"

	"gorm.io/gorm"
)

type taskRepository struct {
	// taskList     []models.Task
	tasksMapListCounter struct {
		sync.RWMutex
		tasksMapList map[int]models.Task
	}
}

func NewTaskRepository(mainDbConn *gorm.DB) domain.TaskRepository {
	// var tasksList []models.Task
	var tasksCounter = struct {
		sync.RWMutex
		tasksMapList map[int]models.Task
	}{tasksMapList: make(map[int]models.Task)}

	defaultTask := models.Task{
		ID:          1,
		Title:       "Task 1",
		Description: "Default Task",
		Status:      "In Progress",
	}
	// tasksList = append(tasksList, defaultTask)
	tasksCounter.Lock()
	tasksCounter.tasksMapList[1] = defaultTask
	tasksCounter.Unlock()
	return &taskRepository{
		// taskList:     tasksList,
		tasksMapListCounter: tasksCounter,
	}
}
