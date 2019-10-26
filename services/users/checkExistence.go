package userservices

import (
	"context"
	"graphql_api/libs/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	Checks if a certain document in the user collection exists.
*/
func CheckExistence(filter bson.M) (bool, error) {
	collection := mongodb.GetClient().Collection("users")
	count, err := collection.CountDocuments(context.TODO(), filter)

	return count > 0, err
}
