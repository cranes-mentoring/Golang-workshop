package repository

import (
	"context"

	"todo-service/pkg/model"
)

func (r *MongoRepository) AddTask(task model.NewTask) error {
	_, err := r.collection.InsertOne(context.TODO(), task)

	return err
}
