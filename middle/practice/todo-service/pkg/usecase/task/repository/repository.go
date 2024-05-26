package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todo-service/pkg/model"
)

type Repository interface {
	AddTask(task model.NewTask) error
	UpdateTask(ID primitive.ObjectID, task model.MongoTask) error
	CompleteTask(ID string) error
	GetAllTasks() ([]model.MongoTask, error)
}
