package handler

import (
	"FirstCrud/internal/middleware"
	"github.com/labstack/echo/v4"
)

// RoutePerson Init routes for Person
func RoutePerson(e *echo.Echo, storage Storage) {
	handler := newPerson(storage)

	// Create a group for all routes
	person := e.Group("/v1/persons")
	person.Use(middleware.Authentication)

	// Register Routes
	person.GET("", handler.GetAll)
	person.POST("", handler.Create)
	person.PUT("/:id", handler.Update)
	person.DELETE("/:id", handler.Delete)
	person.GET("/:id", handler.GetByID)
}

// RouteLogin Init routes for Login
func RouteLogin(e *echo.Echo, storage Storage) {
	login := newLogin(storage)
	e.POST("/v1/login", login.login)
}
