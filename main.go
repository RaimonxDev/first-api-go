package main

import (
	"FirstCrud/internal/auth"
	"FirstCrud/internal/handler"
	"FirstCrud/internal/storage"
	"log"
	"net/http"
)

func main() {
	err := auth.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatal("Can't load certificates")
	}

	store := storage.NewMemory()
	mux := http.NewServeMux()
	//Init routes and storage
	handler.RoutePerson(mux, &store)
	// register login route
	handler.RouteLogin(mux, &store)
	log.Printf("Server running on port %s", ":8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
	// Example of how to use the middleware
	//execute("John", middleware.Log(middleware.Greeter))
	//execute("John", middleware.Log(middleware.Bye))
}

//func execute(name string, f func(string)) {
//	f(name)
//}
