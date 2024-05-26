package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"todo-service/pkg/model"
)

func (r *MongoRepository) GetAllTasks() ([]model.MongoTask, error) {
	var tasks []model.MongoTask
	cursor, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task model.MongoTask
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
