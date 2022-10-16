package main

import (
	"FirstCrud/internal/auth"
	"FirstCrud/internal/handler"
	"FirstCrud/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	err := auth.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatal("Can't load certificates")
	}
	//Init routes and storage
	store := storage.NewMemory()

	//Init Echo
	e := echo.New()
	//Middleware Globals
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler.RouteLogin(e, &store)
	handler.RoutePerson(e, &store)
	log.Printf("Server running on port %s", ":8080")

	// Up server
	err = e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
