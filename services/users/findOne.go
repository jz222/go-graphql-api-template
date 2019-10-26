package userservices

import (
	"context"
	"graphql_api/libs/mongodb"
	"graphql_api/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	Returns one document of the users collection
	that matches the given filter.
*/
func FindOne(filter bson.M) (models.User, error) {
	var user models.User

	collection := mongodb.GetClient().Collection("users")
	result := collection.FindOne(context.TODO(), filter)

	if result.Err() != nil {
		return user, result.Err()
	}

	if err := result.Decode(&user); err != nil {
		log.Println("Failed to decode user with error:", err)
		return user, err
	}

	return user, nil
}
