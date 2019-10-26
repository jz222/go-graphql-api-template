package models

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string
}

type JwtPayload struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
