package repository

import (
	"errors"
	"golang-fiber-crud/helper"
	"golang-fiber-crud/model"
	"time"

	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	DB *gorm.DB
}

func NewTaskRepositoryImpl(DB *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{DB: DB}
}

func (t *TaskRepositoryImpl) Delete(taskId int) {
	var task model.Task
	result := t.DB.Where("id=?", taskId).Delete(task)
	helper.ErrorPanic(result.Error)
}

func (t *TaskRepositoryImpl) FindAll() []model.Task {
	var tasks []model.Task

	result := t.DB.Preload("items").Find(&tasks)
	helper.ErrorPanic(result.Error)
	return tasks
}

func (t *TaskRepositoryImpl) FindById(taskId int) (model.Task, error) {
	var task model.Task
	result := t.DB.Preload("Items").Find(&task, taskId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return task, errors.New("Task not found")
		} else {
			panic(result.Error)
		}
	}
	return task, nil
}

func (t *TaskRepositoryImpl) Save(task model.Task) {
	task.CreatedAt = time.Now().Unix()
	task.ModifiedAt = time.Now().Unix()
	result := t.DB.Create(&task)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (t *TaskRepositoryImpl) Update(task model.Task) {
	task.ModifiedAt = time.Now().Unix()
	result := t.DB.Save(&task)
	if result.Error != nil {
		panic(result.Error)
	}
}
