package service

import (
	"todo-service/pkg/mapper"
	"todo-service/pkg/model"
)

func (s *Service) GetAllTasks() ([]model.Task, error) {
	tasks, err := s.Repo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	return mapTo(tasks), nil
}

func mapTo(list []model.MongoTask) []model.Task {
	mappedTasks := make([]model.Task, 0)
	for _, task := range list {
		mappedTasks = append(mappedTasks, mapper.MapToModel(task))
	}

	return mappedTasks
}
