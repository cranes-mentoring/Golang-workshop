package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"todo-service/pkg/model"
)

func (r *MongoRepository) UpdateTask(task model.MongoTask) error {
	objID, err := primitive.ObjectIDFromHex(task)
	if err != nil {
		return err
	}

	filter := bson.D{{"id", task.ID}}
	update := bson.D{{"$set", bson.D{{"title", task.Title}, {"completed", task.Completed}}}}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}
