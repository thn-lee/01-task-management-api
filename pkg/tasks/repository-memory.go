package tasks

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/thn-lee/01-task-management-api/pkg/domain"
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
	"github.com/thn-lee/01-task-management-api/pkg/utils"

	// utils "github.com/zercle/gofiber-utils"

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

func (r *taskRepository) GetTask(taskID uint) (task models.Task, err error) {
	for _, task := range r.taskList {
		if task.ID == taskID {
			return task, nil
		}
	}

	log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
	err = fiber.NewError(http.StatusInternalServerError, err.Error())
	return
}

func (r *taskRepository) GetTasks(criteria models.Task) (tasks []models.Task, err error) {
	tasks = r.taskList
	return
}

func (r *taskRepository) CreateTask(task *models.Task) (result models.Task, err error) {
	task.ID = uint(len(r.taskList) + 1)
	fmt.Println(">>>> passed 1")
	task.CreatedAt = time.Now()
	fmt.Println(">>>> passed 2")
	task.UpdatedAt = time.Now()
	fmt.Println(">>>> passed 3")
	r.taskList = append(r.taskList, *task)
	fmt.Println(">>>> passed 4")

	result = *task
	fmt.Println(">>>> passed 5")
	return
}

func (r *taskRepository) EditTask(taskID uint, incomingTask models.Task) (result models.Task, err error) {
	for taskIndex, task := range r.taskList {
		if task.ID == taskID {
			// update title
			if incomingTask.Title != "" {
				r.taskList[taskIndex].Title = incomingTask.Title
			}
			// update title description
			if incomingTask.Description != "" {
				r.taskList[taskIndex].Description = incomingTask.Description
			}
			r.taskList[taskIndex].UpdatedAt = time.Now()
			fmt.Println(">>>>>>> updated task")
			return r.taskList[taskIndex], nil
		}
	}

	// task not found
	log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
	err = fiber.NewError(http.StatusInternalServerError, err.Error())
	return
}

func (r *taskRepository) EditTaskStatus(taskID uint, status string) (result models.Task, err error) {
	for taskIndex, task := range r.taskList {
		if task.ID == taskID {
			// update title
			if status != "" {
				r.taskList[taskIndex].Status = status
			}
			r.taskList[taskIndex].UpdatedAt = time.Now()
			fmt.Println(">>>>>>> updated task")
			return r.taskList[taskIndex], nil
		}
	}

	// task not found
	log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
	err = fiber.NewError(http.StatusInternalServerError, err.Error())
	return
}

func (r *taskRepository) DeleteTask(taskID uint) (err error) {
	for taskIndex, task := range r.taskList {
		if task.ID == taskID {
			r.taskList = append(r.taskList[:taskIndex], r.taskList[taskIndex+1:]...)
			return nil
		}
	}

	// task not found
	log.Printf("source: %+v\nerr: %+v", utils.WhereAmI(), err)
	err = fiber.NewError(http.StatusInternalServerError, err.Error())
	return
}
