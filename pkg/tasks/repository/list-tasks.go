package tasks

import (
	models "github.com/thn-lee/01-task-management-api/pkg/models/api"
)

func (r *taskRepository) ListTasks(criteria models.Task) (tasks map[int]models.Task, err error) {
	r.tasksMapListCounter.RLock()
	defer r.tasksMapListCounter.RUnlock()
	tasks = r.tasksMapListCounter.tasksMapList

	// if criteria.Title != "" {

	// 	// for index, task := range r.tasksMapListCounter.tasksMapList {
	// 	// 	if task.Title == criteria.Title {
	// 	// 		r.tasksMapListCounter[index]
	// 	// 	}
	// 	// }
	// }
	// if criteria.Description != "" {
	// }
	// if criteria.Status != "" {
	// }

	return
}
