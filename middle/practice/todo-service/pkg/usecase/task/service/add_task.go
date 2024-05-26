package service

import (
	"todo-service/pkg/mapper"
	"todo-service/pkg/model"
)

func (s *Service) AddTask(task model.Task) error {
	return s.Repo.AddTask(mapper.MapToNew(task))
}
