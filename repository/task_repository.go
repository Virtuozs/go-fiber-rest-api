package repository

import "golang-fiber-crud/model"

type TaskRepository interface {
	Save(task model.Task)
	Update(task model.Task)
	Delete(taskId int)
	FindById(taskId int) (model.Task, error)
	FindAll() []model.Task
}
