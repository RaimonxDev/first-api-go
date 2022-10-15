package model

import "github.com/golang-jwt/jwt/v4"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
