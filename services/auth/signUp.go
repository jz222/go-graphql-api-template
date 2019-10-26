package authservices

import (
	"errors"
	"graphql_api/models"
	userservices "graphql_api/services/users"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	Checks if the given email address was used before and stores
	the new user in the database with the hashed password.
*/
func SignUp(newUser models.User) (string, error) {
	userExists, err := userservices.CheckExistence(bson.M{"email": newUser.Email})

	if err != nil {
		return "", err
	}

	if userExists {
		return "", errors.New("User with the given email address already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)

	if err != nil {
		return "", err
	}

	newUser.Password = string(hash)

	if _, err := userservices.Create(newUser); err != nil {
		return "", err
	}

	jwt, err := GenerateJwt(string(newUser.ID.Hex()))

	if err != nil {
		return "", err
	}

	return jwt, nil
}
