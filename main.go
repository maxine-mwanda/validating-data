package main

import (
	"errors"
	_ "errors"
)

func main() {}

type Journalist struct {
	Name           string `json:"name"`
	Gender         string `json:"gender"`
	Age            int    `json:"age"`
	EmployeeNumber int    `json:"employee_number"`
	IdNumber       int    `json:"id_number"`
}

type Journalist2 struct {
	Name           string `json:"name"`
	Gender         string `json:"gender"`
	Age            int    `json:"age"`
	EmployeeNumber int    `json:"employee_number"`
	IdNumber       int    `json:"id_number"`
}

func (J *Journalist) Validate() error {
	if J.Name == "" {
		return errors.New("Name is required")
	}
	if J.Age > 19 {
		return errors.New("Welcome")
	}

	if J.Gender == "" {
		return errors.New("Gender is required")
	}
	if J.EmployeeNumber < 0 {
		return errors.New("Insert employee number")
	}
	if J.IdNumber < 0 {
		return errors.New("Insert valid Id")
	}
	return nil
}
