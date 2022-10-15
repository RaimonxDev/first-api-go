package handler

import (
	"FirstCrud/internal/auth"
	"FirstCrud/internal/model"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {

	data := model.Login{}
	err := c.Bind(&data) // Bind data from request to model
	if err != nil {
		// Return error like json
		return c.JSON(http.StatusBadRequest, errors.New("invalid json"))
	}

	if !isLoginValid(data) {
		response := NewResponse("Email or Password not valid", http.StatusBadRequest, "Bad Request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	token, err := auth.GenerateToken(&data)
	if err != nil {
		response := NewResponse("Error to generate token", http.StatusInternalServerError, "Internal Server Error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	dataToken := map[string]string{
		"token": token,
		"user":  data.Email,
	}
	response := NewResponse("Login OK", http.StatusOK, nil, dataToken)
	return c.JSON(http.StatusOK, response)
}

// Simulate a login
func isLoginValid(data model.Login) bool {
	if data.Email == "" {
		return false
	}
	if data.Password == "" {
		return false
	}
	return data.Email == "ramon@gmail.com" && data.Password == "123456"
}
