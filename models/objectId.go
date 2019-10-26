package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoDbDocument struct {
	ID *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
}
