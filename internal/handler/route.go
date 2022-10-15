package handler

import (
	"FirstCrud/internal/middleware"
	"net/http"
)

func RoutePerson(mux *http.ServeMux, storage Storage) {
	person := newPerson(storage)
	mux.HandleFunc("/v1/persons", middleware.Authentication(person.GetAll))
	mux.HandleFunc("/v1/persons/create", middleware.Authentication(person.Create))
	mux.HandleFunc("/v1/persons/update", middleware.Log(person.Update))
	mux.HandleFunc("/v1/persons/delete", middleware.Log(person.Delete))
	mux.HandleFunc("/v1/persons/person", middleware.Log(person.GetByID))
}

// Route Login
func RouteLogin(mux *http.ServeMux, storage Storage) {
	login := newLogin(storage)
	mux.HandleFunc("/v1/login", login.login)
}
