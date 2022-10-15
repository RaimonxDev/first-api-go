package middleware

import (
	"FirstCrud/internal/auth"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		defer timeTrack(time.Now(), r.URL.Path)
		log.Printf("Request: %s, method: %s", r.URL.Path, r.Method)
	}
}
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// Authentication Middleware
func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not Authorized")
		}
		_, err := auth.ValidateToken(token)

		if err != nil {
			return c.JSON(http.StatusForbidden, "Forbidden")
		}

		return f(c)
	}
}
