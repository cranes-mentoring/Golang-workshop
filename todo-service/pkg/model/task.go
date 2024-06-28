package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoTask struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title"`
	Completed bool               `json:"completed"`
}
