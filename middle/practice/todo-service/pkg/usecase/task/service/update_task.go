package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todo-service/pkg/mapper"
	"todo-service/pkg/model"
)

func (s *Service) UpdateTask(ID string, task model.Task) error {
	premID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	return s.Repo.UpdateTask(premID, mapper.MapToDto(task))
}
