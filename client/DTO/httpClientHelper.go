package DTO

import (
	"io"
	"log"
	"net/http"
)

func HttpClientHelper(method, url, token string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Error to create a new request: %v", err)
	}
	// Set header to request ALWAYS
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error to do request: %v", err)
	}
	return resp
}
