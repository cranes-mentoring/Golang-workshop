package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todo-service/pkg/model"
)

func (r *MongoRepository) UpdateTask(ID primitive.ObjectID, task model.MongoTask) error {
	filter := bson.D{{"_id", ID}}
	update := bson.D{{"$set", bson.D{{"title", task.Title}, {"completed", task.Completed}}}}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)

	return err
}
