package storage

import (
	"FirstCrud/internal/model"
	"fmt"
)

type Memory struct {
	CurrentId int
	Persons   map[int]model.Person
}

// NewMemory Return one instances memory
func NewMemory() Memory {
	persons := make(map[int]model.Person)
	return Memory{
		CurrentId: 0,
		Persons:   persons,
	}
}

// Create Person
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonNotBeNil
	}
	m.CurrentId++
	m.Persons[m.CurrentId] = *person
	return nil
}

// Update memory
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonNotBeNil
	}
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesExists)
	}
	m.Persons[ID] = *person

	return nil
}

// Delete person
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesExists)
	}
	delete(m.Persons, ID)
	return nil
}

// GetByID person by ID
func (m *Memory) GetByID(ID int) (model.Person, error) {

	if _, ok := m.Persons[ID]; !ok {
		return model.Person{}, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesExists)
	}
	return m.Persons[ID], nil
}

// GetAll persons
func (m *Memory) GetAll() (model.Persons, error) {
	persons := make(model.Persons, 0)
	for _, person := range m.Persons {
		persons = append(persons, person)
	}
	return persons, nil
}
