package repository

import (
	"todo-service/pkg/model"
)

type Repository interface {
	AddTask(task model.NewTask) error
	CompleteTask(ID string) error
}
