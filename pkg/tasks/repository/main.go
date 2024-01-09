package tasks

import (
	"github.com/thn-lee/01-task-management-api/pkg/domain"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"

	"gorm.io/gorm"
)

type taskRepository struct {
	taskList []models.Task
}

func NewTaskRepository(mainDbConn *gorm.DB) domain.TaskRepository {
	var tasksList []models.Task
	tasksList = append(tasksList, models.Task{
		ID:          1,
		Title:       "Task 1",
		Description: "Default Task",
		Status:      "In Progress",
	})
	return &taskRepository{
		taskList: tasksList,
	}
}
