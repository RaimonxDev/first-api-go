package DTO

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginClient(url, email, password string) LoginResponse {
	var err error

	login := Login{
		Email:    email,    // Set email from request and parse to json
		Password: password, // Set password from request and parse to json
	}
	data := bytes.NewBuffer([]byte{}) // IMPORTANT Create a buffer to store data
	// Encode data to json and store in buffer data
	err = json.NewEncoder(data).Encode(&login) // Encode data to json

	if err != nil {
		log.Fatalf("Error to encode data to json: %v", err)
	}
	// Create a request to url with method POST and body data
	resp := HttpClientHelper(http.MethodPost, url, "", data)
	defer resp.Body.Close() // Close body response ALWAYS

	// NEED USED io.ReadAll to read body response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error to read body: %v", err)
	}

	// Only status code 200 is valid
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error not status OK : %v, body: %s", resp.StatusCode, string(body))
	}

	dataResponse := LoginResponse{} // Create a struct to store data response

	//Decoder data response to struct dataResponse
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse) // Encode body(resp.body) and save to dataResponse
	if err != nil {
		log.Fatalf("Error to decode data: %v", err)
	}
	return dataResponse

}

type GeneralResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
}

// LoginResponse Create Data struct, To save Response from server
type LoginResponse struct {
	GeneralResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}
