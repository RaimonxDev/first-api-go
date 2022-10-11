package handler

import (
	"FirstCrud/internal/model"
	"encoding/json"
	"errors"
	http "net/http"
	"strconv"
)

// Person implements the handler interface
type person struct {
	storage Storage
}

// NewPerson return one instance of person
func newPerson(storage Storage) person {
	return person{
		storage: storage,
	}
}

func (p *person) Create(w http.ResponseWriter, r *http.Request) {

	// Only POST petition
	if r.Method != http.MethodPost {
		response := NewResponse("Only POST petition", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}

	data := model.Person{}
	// Decode the body
	// If the body is not a json, return an error
	// NewDecoder parse a json and validate the syntax in GO
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := NewResponse("JSON NO VALID", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	err = p.storage.Create(&data)
	if err != nil {
		response := NewResponse("Error to create person", http.StatusInternalServerError, "Internal Server Error", nil)
		response.ToJSON(w)
		return
	}

	response := NewResponse("Person created", http.StatusCreated, nil, data)
	response.ToJSON(w)
}

func (p *person) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := NewResponse("Only GET petition", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	// call method GetAll from storage or database
	persons, err := p.storage.GetAll()
	if err != nil {
		response := NewResponse("Error to get persons", http.StatusInternalServerError, "Internal Server Error", nil)
		response.ToJSON(w)
		return
	}
	// Encode the body
	response := NewResponse("Success", http.StatusOK, nil, persons)
	err = response.ToJSON(w)
	if err != nil {
		response := NewResponse("Error to encode persons", http.StatusInternalServerError, "Internal Server Error", nil)
		response.ToJSON(w)
		return
	}
}

func (p *person) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		response := NewResponse("Only PUT petition", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}

	data := model.Person{}
	// Convert the string to int with func strconv.Atoi
	// Get the id from the url query
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := NewResponse("Error to convert id", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	// Decode the body
	// If the body is not a json, return an error
	// NewDecoder parse a json and validate the syntax in GO
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := NewResponse("JSON NO VALID", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := NewResponse("Error to update person", http.StatusInternalServerError, "Internal Server Error", nil)
		response.ToJSON(w)
		return
	}

	response := NewResponse("Person updated", http.StatusOK, nil, data)
	err = response.ToJSON(w)
	if err != nil {
		response := NewResponse("Error to encode persons", http.StatusInternalServerError, "Internal Server Error", nil)
		response.ToJSON(w)
		return
	}

}

func (p *person) Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		response := NewResponse("Only DELETE petition", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	//required ID
	if r.URL.Query().Get("id") == "" {
		response := NewResponse("ID is required", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	// Convert the string to int with func strconv.Atoi
	// Get the id from the url query
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := NewResponse("Error to convert id", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}

	err = p.storage.Delete(ID)
	// When the id is not found, return an error,
	// Used the errors package from GO to compare the error with the error from the database
	if errors.Is(err, model.ErrIDPersonDoesExists) {
		response := NewResponse("ID does not exists", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	if err != nil {
		response := NewResponse("Error to delete person", http.StatusInternalServerError, "Internal Server Error", nil)
		response.ToJSON(w)
		return
	}
	response := NewResponse("Person deleted", http.StatusOK, nil, nil)
	err = response.ToJSON(w)

}

func (p *person) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := NewResponse("Only GET petition", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	//required ID
	if r.URL.Query().Get("id") == "" {
		response := NewResponse("ID is required", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	// Convert the string to int with func strconv.Atoi
	// Get the id from the url query
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := NewResponse("ID not valid", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	// Get User by ID
	person, err := p.storage.GetByID(ID)

	if errors.Is(err, model.ErrIDPersonDoesExists) {
		response := NewResponse("ID does not exists", http.StatusBadRequest, "Bad Request", nil)
		response.ToJSON(w)
		return
	}
	if err != nil {
		response := NewResponse("Error to get person", http.StatusInternalServerError, "Internal Server Error", nil)
		response.ToJSON(w)
		return
	}
	response := NewResponse("Success", http.StatusOK, nil, person)
	err = response.ToJSON(w)
}
