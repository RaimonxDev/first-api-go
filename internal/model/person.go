package model

type Communities []Community
type Community struct {
	Name string `json:"name"`
}

type Persons []Person

type Person struct {
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Communities Communities `json:"communities"`
}
