package main

import (
	"FirstCrud/client/DTO"
	"FirstCrud/internal/model"
	"fmt"
)

var (
	url = "http://localhost:8080"
)

func main() {
	loginClient := DTO.LoginClient(url+"/v1/login", "ramon@gmail.com", "123456")
	fmt.Println(loginClient)

	person := model.Person{
		Name:        "Ramon",
		Age:         30,
		Communities: []model.Community{{Name: "Golang"}, {Name: "Python"}},
	}
	personUPDATED := model.Person{
		Name:        "JAVIER",
		Age:         30,
		Communities: []model.Community{{Name: "Golang"}, {Name: "Python"}},
	}

	DTO.CreatePerson(url+"/v1/persons", loginClient.Data.Token, &person)

	DTO.UpdatePerson(url+"/v1/persons/", loginClient.Data.Token, "8", &personUPDATED)

}
