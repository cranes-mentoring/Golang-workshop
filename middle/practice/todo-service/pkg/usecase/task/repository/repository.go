package repository

import (
	"todo-service/pkg/model"
)

type Repository interface {
	AddTask(task model.NewTask) error
	UpdateTask(task model.MongoTask) error
	CompleteTask(ID string) error
	GetAllTasks() ([]model.MongoTask, error)
}
