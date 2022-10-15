package handler

import (
	"FirstCrud/internal/model"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
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

func (p *person) Create(e echo.Context) error {
	data := model.Person{}
	// Decode the body
	err := e.Bind(&data)

	if err != nil {
		response := NewResponse("JSON NO VALID", http.StatusBadRequest, "Bad Request", nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	// execute the method Create from storage
	err = p.storage.Create(&data)
	if err != nil {
		response := NewResponse("Error to create person", http.StatusInternalServerError, "Internal Server Error", nil)
		return e.JSON(http.StatusInternalServerError, response)
	}

	response := NewResponse("Person created", http.StatusCreated, nil, data)
	return e.JSON(http.StatusCreated, response)
}

func (p *person) GetAll(e echo.Context) error {
	// call method GetAll from storage or database
	persons, err := p.storage.GetAll()
	if err != nil {
		response := NewResponse("Error to get persons", http.StatusInternalServerError, "Internal Server Error", nil)
		return e.JSON(http.StatusInternalServerError, response)
	}
	// Encode the body
	response := NewResponse("Success", http.StatusOK, nil, persons)
	return e.JSON(http.StatusOK, response)
}

func (p *person) Update(e echo.Context) error {

	var err error
	// required ID
	if err := e.Param("id"); err == "" {
		return e.JSON(http.StatusBadRequest, "ID is required")
	}
	id, _ := strconv.Atoi(e.Param("id"))

	// Decode the body
	data := model.Person{}
	err = e.Bind(&data)

	if err != nil {
		return e.JSON(http.StatusBadRequest, "JSON NO VALID")
	}

	if err = p.storage.Update(id, &data); errors.Is(err, model.ErrIDPersonDoesExists) {
		return e.JSON(http.StatusBadRequest, "ID does not exists")
	}

	response := NewResponse("Person updated", http.StatusOK, nil, data)
	return e.JSON(http.StatusOK, response)

}

func (p *person) Delete(e echo.Context) error {
	var err error
	//required ID
	if err := e.Param("id"); err == "" {
		return e.JSON(http.StatusBadRequest, "ID is required")
	}

	id, _ := strconv.Atoi(e.Param("id"))

	err = p.storage.Delete(id)
	// When the id is not found, return an error,
	// Used the errors package from GO to compare the error with the error from the database
	if errors.Is(err, model.ErrIDPersonDoesExists) {
		response := NewResponse("ID does not exists", http.StatusBadRequest, "Bad Request", nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := NewResponse("Error to delete person", http.StatusInternalServerError, "Internal Server Error", nil)
		return e.JSON(http.StatusInternalServerError, response)
	}
	response := NewResponse("Person deleted", http.StatusOK, nil, nil)
	return e.JSON(http.StatusOK, response)

}

func (p *person) GetByID(e echo.Context) error {

	//required ID
	if err := e.Param("id"); err == "" {
		return e.JSON(http.StatusBadRequest, "ID is required")
	}
	id, _ := strconv.Atoi(e.Param("id"))

	// Get User by ID
	person, err := p.storage.GetByID(id)

	if errors.Is(err, model.ErrIDPersonDoesExists) {
		response := NewResponse("ID does not exists", http.StatusBadRequest, "Bad Request", nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := NewResponse("Error to get person", http.StatusInternalServerError, "Internal Server Error", nil)
		return e.JSON(http.StatusInternalServerError, response)
	}
	response := NewResponse("Success", http.StatusOK, nil, person)
	return e.JSON(http.StatusOK, response)
}
