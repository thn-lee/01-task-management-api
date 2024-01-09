package tasks

// import (
// 	"fmt"

// 	"github.com/thn-lee/01-task-management-api/pkg/domain"
// 	models "github.com/thn-lee/01-task-management-api/pkg/models/api"

// 	helpers "github.com/thn-lee/01-task-management-api/pkg/utils"
// 	"gorm.io/gorm"
// )

// type taskRepository struct {
// 	MainDbConn *gorm.DB
// }

// func NewTaskRepository(mainDbConn *gorm.DB) domain.TaskRepository {
// 	return &taskRepository{
// 		MainDbConn: mainDbConn,
// 	}
// }

// func (r *taskRepository) GetTask(taskID uint) (task models.Task, err error) {
// 	if r.MainDbConn == nil {
// 		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
// 		return
// 	}

// 	dbTx := r.MainDbConn.Model(&models.Task{})
// 	dbTx = dbTx.Where(models.Task{ID: taskID})
// 	err = dbTx.Take(&task).Error

// 	return
// }

// func (r *taskRepository) ListTasks(criteria models.Task) (tasks []models.Task, err error) {
// 	if r.MainDbConn == nil {
// 		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
// 		return
// 	}

// 	dbTx := r.MainDbConn.Model(&models.Task{})

// 	if len(criteria.ID) != 0 {
// 		dbTx = dbTx.Where(models.Task{ID: criteria.ID})
// 	} else {
// 		// if len(criteria.FullName) != 0 {
// 		// 	dbTx = dbTx.Where("title LIKE ?", "%"+criteria.FullName+"%")
// 		// }
// 	}

// 	err = dbTx.Find(&tasks).Error

// 	return
// }

// func (r *taskRepository) CreateTask(task models.TaskBody) (result models.Task, err error) {
// 	if r.MainDbConn == nil {
// 		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
// 		return
// 	}

// 	dbTx := r.MainDbConn.Begin()
// 	defer dbTx.Rollback()

// 	dbTx = dbTx.Model(&models.Task{})

// 	if err = dbTx.Create(task).Error; err != nil {
// 		return
// 	}

// 	err = dbTx.Commit().Error

// 	return
// }

// func (r *taskRepository) EditTask(taskID uint, requestTask models.TaskBody) (result models.Task, err error) {
// 	if r.MainDbConn == nil {
// 		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
// 		return
// 	}

// 	dbTx := r.MainDbConn.Begin()
// 	defer dbTx.Rollback()

// 	dbTx = dbTx.Model(&models.Task{})
// 	dbTx = dbTx.Where(models.Task{ID: taskID})

// 	if err = dbTx.Updates(task).Error; err != nil {
// 		return
// 	}

// 	err = dbTx.Commit().Error

// 	return
// }

// func (r *taskRepository) DeleteTask(taskID uint) (err error) {
// 	if r.MainDbConn == nil {
// 		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
// 		return
// 	}

// 	dbTx := r.MainDbConn.Begin()
// 	defer dbTx.Rollback()

// 	dbTx = dbTx.Model(&models.Task{})
// 	dbTx = dbTx.Where(models.Task{ID: taskID})

// 	if err = dbTx.Delete(&models.Task{}).Error; err != nil {
// 		return
// 	}

// 	err = dbTx.Commit().Error

// 	return
// }
