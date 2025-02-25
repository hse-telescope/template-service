package server

import (
	"github.com/hse-telesope/template-service/internal/providers/people"
)

type Person struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func ServerPerson2ProviderPerson(person Person) people.Person {
	return people.Person{
		ID:       person.ID,
		Username: person.Username,
	}
}

func ProviderPerson2ServerPerson(person people.Person) Person {
	return Person{
		ID:       person.ID,
		Username: person.Username,
	}
}
