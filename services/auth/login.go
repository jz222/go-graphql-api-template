package authservices

import (
	"errors"
	userservices "graphql_api/services/users"
	"graphql_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

/*
	Checks if the user with the given email exists and
	verifies the given password.
*/
func Login(credentials models.Credentials) (string, error) {
	user, err := userservices.FindOne(bson.M{"email": credentials.Email})

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("Incorrect password")
	}

	jwt, err := GenerateJwt(string(user.ID.Hex()))

	if err != nil {
		return "", err
	}

	return jwt, nil
}
