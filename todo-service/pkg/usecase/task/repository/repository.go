package repository

import (
	"context"

	"todo-service/pkg/model"
)

type Repository interface {
	AddTask(ctx context.Context, task model.MongoTask) error
	CompleteTask(ctx context.Context, ID string) error
}
