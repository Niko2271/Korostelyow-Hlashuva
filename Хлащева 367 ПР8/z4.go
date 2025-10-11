package main

import (
	"fmt"
	"strings"
)

type Validator interface {
	Validate(string) error
}

type Field struct {
	Name      string
	Value     string
	Validators []Validator
}

type Required struct{}

func (r Required) Validate(v string) error {
	if v == "" {
		return fmt.Errorf("нельзя пусто")
	}
	return nil
}

type MinLimsim struct {
	Length int
}

func (m MinLimsim) Validate(v string) error {
	if len(v) < m.Length {
		return fmt.Errorf("минимум %d символов", m.Length)
	}
	return nil
}

type Email struct{}

func (e Email) Validate(v string) error {
	if !strings.Contains(v, "@") {
		return fmt.Errorf("это не почта")
	}
	return nil
}

func (f *Field) Validate() []error {
	var errors []error
	for _, validator := range f.Validators {
		if err := validator.Validate(f.Value); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func main() {
	field := Field{
		Name:  "email",
		Value: "t",
		Validators: []Validator{
			Required{},
			MinLimsim{Length: 3},
			Email{},
		},
	}
	
	for _, err := range field.Validate() {
		fmt.Println(err)
	}
}
