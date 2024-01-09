package tasks

func (u *taskUsecase) DeleteTask(taskID uint) (err error) {
	return u.taskRepo.DeleteTask(taskID)
}
