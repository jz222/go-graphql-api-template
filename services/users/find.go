package userservices

import (
	"context"
	"graphql_api/libs/mongodb"
	"graphql_api/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	Finds and returns all documents of the users collection
	that match the given filter.
*/
func Find(filter bson.M) ([]*models.User, error) {
	var users []*models.User

	collection := mongodb.GetClient().Collection("users")
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal("Failed to query user collection with error:", err)
		return users, err
	}

	for cur.Next(context.TODO()) {
		var user models.User

		err = cur.Decode(&user)

		if err != nil {
			log.Fatal("Failed to decode user with error:", err)
			return users, err
		}

		users = append(users, &user)
	}

	return users, nil
}
