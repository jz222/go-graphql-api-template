package middlewares

import (
	"context"
	"fmt"
	"graphql_api/keys"
	"graphql_api/models"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var CtxUserFromJwt = models.ContextKey{"user"}

func setEmptyContext(next http.Handler, w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), CtxUserFromJwt, nil)
	r = r.WithContext(ctx)
	next.ServeHTTP(w, r)
}

/*
	Middleware to verify JWT's sent as Authorization header.
*/
func VerifyJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authenticationHeader := r.Header.Get("Authorization")
		splitHeader := strings.Split(authenticationHeader, " ")

		if len(splitHeader) != 2 {
			setEmptyContext(next, w, r)
			return
		}

		token, err := jwt.ParseWithClaims(splitHeader[1], &models.JwtPayload{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("An error occured while verifying the token")
			}

			return []byte(keys.GetKeys().SECRET), nil
		})

		if err != nil || !token.Valid {
			setEmptyContext(next, w, r)
			return
		}

		claims, ok := token.Claims.(*models.JwtPayload)

		if !ok {
			setEmptyContext(next, w, r)
			return
		}

		ctx := context.WithValue(r.Context(), CtxUserFromJwt, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
