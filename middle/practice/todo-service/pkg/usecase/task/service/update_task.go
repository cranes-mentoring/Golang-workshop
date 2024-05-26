package service

import (
	"todo-service/pkg/mapper"
	"todo-service/pkg/model"
)

func (s *Service) UpdateTask(task model.Task) error {
	return s.Repo.UpdateTask(mapper.MapToDto(task))
}
