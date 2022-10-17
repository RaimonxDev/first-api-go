package DTO

import (
	model "FirstCrud/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreatePerson(url, token string, person *model.Person) GeneralResponse {
	var err error
	data := bytes.NewBuffer([]byte{}) // IMPORTANT Create a buffer to store data
	// Encode data to json and store in buffer data
	err = json.NewEncoder(data).Encode(&person) // Encode data to json

	if err != nil {
		log.Fatalf("Error to encode data to json: %v", err)
	}
	// Create a request to url with method POST , token, body and data
	resp := HttpClientHelper(http.MethodPost, url, token, data)
	defer resp.Body.Close() // Close body response ALWAYS, free memory

	// NEED USED io.ReadAll to read body response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error to read body: %v", err)
	}

	// Only status code 200 is valid
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Error : %v, body: %s", resp.StatusCode, string(body))
	}

	dataResponse := GeneralResponse{} // Create a struct to store data response

	//Decoder data response to struct dataResponse
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse) // Encode body(resp.body) and save to dataResponse
	if err != nil {
		log.Fatalf("Error to decode data: %v", err)
	}

	fmt.Println(string(body))
	return dataResponse

}

func UpdatePerson(url, token, id string, person *model.Person) {
	var err error
	data := bytes.NewBuffer([]byte{}) // IMPORTANT Create a buffer to store data
	// Encode data to json and store in buffer data
	err = json.NewEncoder(data).Encode(&person) // Encode data to json

	if err != nil {
		log.Fatalf("Error to encode data to json: %v", err)
	}
	// Create a request to url with method PUT , token, body and data
	resp := HttpClientHelper(http.MethodPut, url+id, token, data)
	defer resp.Body.Close() // Close body response ALWAYS, free memory

	// NEED USED io.ReadAll to read body response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error to read body: %v", err)
	}

	// Only status code 200 is valid
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error : %v, body: %s", resp.StatusCode, string(body))
	}

	fmt.Println(string(body))
}
