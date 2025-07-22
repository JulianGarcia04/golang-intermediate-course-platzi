package poo

import (
	"fmt"
	"log"
)

// Golang no tiene herencia como otros lenguajes de programación
/*
Golang maneja la composición sobre la herencia, es decir que permite
que cada struct tenga sus propiedades independientes, pero yo puedo crear un struct apartir de otros como si
esos otros otros fueran componentes
*/

type Person struct {
	name string
	age  int
}

type BaseEmployee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	BaseEmployee
}

func (fte FullTimeEmployee) String() string {
	return fmt.Sprintf("Full Time Employee with id %d and name %s and %d years old", fte.id, fte.name, fte.age)
}

func (fte FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

func Inheritance() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.id = 1
	ftEmployee.age = 22
	ftEmployee.name = "Tom"

	log.Println(ftEmployee.String())
}
