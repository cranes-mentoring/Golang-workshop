package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *MongoRepository) CompleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objID}}
	update := bson.D{{"$set", bson.D{{"completed", true}}}}
	_, err = r.collection.UpdateOne(context.TODO(), filter, update)

	return err
}
