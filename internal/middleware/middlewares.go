package middleware

import (
	"log"
	"net/http"
	"time"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		defer timeTrack(time.Now(), r.URL.Path)
		log.Printf("Request: %s, method: %s", r.URL.Path, r.Method)
		Authentication(next)(w, r)
	}
}
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// Authentication Middleware
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// Simulate authentication
		if token != "Bearer 123456789" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

// Middleware is a function that wraps a function and returns a function.

// Middleware Custom type for middleware
//type Middleware func(string)
//NO USED
//func Log(f Middleware) Middleware {
//	return func(s string) {
//		fmt.Println(time.Now().Format("15:04:05")) // Print the current time
//		f(s)
//	}
//}
