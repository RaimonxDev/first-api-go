package auth

import (
	model "FirstCrud/internal/model"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(login *model.Login) (string, error) {

	claim := model.Claims{
		Email: login.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "FirstCrud",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)

	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(t string) (model.Claims, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claims{}, verifyToken)
	if err != nil {
		return model.Claims{}, nil
	}

	if !token.Valid {
		return model.Claims{}, errors.New("Token not valid")
	}

	claim, ok := token.Claims.(*model.Claims)

	if !ok {
		return model.Claims{}, errors.New("Claims no valid")
	}
	return *claim, nil
}

func verifyToken(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
