package authservices

import (
	"graphql_api/keys"
	"graphql_api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
	Generates and returns a JWT.
*/
func GenerateJwt(userId string) (string, error) {
	timestamp := time.Now().Unix()
	expiresAt := timestamp + 600

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtPayload{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  timestamp,
		},
	})

	signedToken, err := token.SignedString([]byte(keys.GetKeys().SECRET))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
