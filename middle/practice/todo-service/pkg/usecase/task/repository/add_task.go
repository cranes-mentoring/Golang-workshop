package repository

import (
	"context"

	"todo-service/pkg/model"
)

func (r *MongoRepository) AddTask(ctx context.Context, task model.NewTask) error {
	_, err := r.collection.InsertOne(ctx, task)

	return err
}
