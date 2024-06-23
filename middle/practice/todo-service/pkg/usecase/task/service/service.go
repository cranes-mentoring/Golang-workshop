package service

import (
	"todo-service/pkg/model"
	"todo-service/pkg/usecase/task/repository"
)

type TodoService interface {
	AddTask(task model.Task) error
	UpdateTask(ID string, task model.Task) error
	CompleteTask(ID string) error
	GetAllTasks() ([]model.Task, error)
}

type Service struct {
	Repo repository.Repository
}

func NewService(repo repository.Repository) TodoService {
	return &Service{Repo: repo}
}
