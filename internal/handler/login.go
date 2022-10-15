package handler

import (
	"FirstCrud/internal/auth"
	"FirstCrud/internal/model"
	"encoding/json"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		response := NewResponse(
			"METHOD NO ALLOWED",
			http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	if r.Header.Get("Authorization") == "" {
		response := NewResponse("MISSING TOKEN", http.StatusBadRequest, "Bad request", nil)
		response.ToJSON(w)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		respose := NewResponse("JSON NO VALID", http.StatusBadRequest, "Bad Request", nil)
		respose.ToJSON(w)
		return
	}

	if !isLoginValid(data) {
		response := NewResponse("Email or Password no valid", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	token, err := auth.GenerateToken(&data)
	if err != nil {
		response := NewResponse("Error to generate token", http.StatusInternalServerError, "Internal Server Error", nil)
		response.ToJSON(w)
		return
	}
	dataToken := map[string]string{
		"token": token,
		"user":  data.Email,
	}
	response := NewResponse("Login OK", http.StatusOK, nil, dataToken)
	response.ToJSON(w)
	return
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
