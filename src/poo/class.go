package poo

import (
	"fmt"
	"log"
)

type Employee struct {
	id   int
	name string
}

func (e *Employee) setId(id int) {
	e.id = id
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) getId() int {
	return e.id
}

func (e *Employee) getName() string {
	return e.name
}

func (e *Employee) String() string {
	return fmt.Sprintf("id: %d, name: %s", e.id, e.name)
}

func Main() {
	e := Employee{}

	log.Printf("%+v", e)

	e.id = 1

	e.name = "Jack"

	log.Printf("%+v", e)

	e.setId(1)
	e.setName("Hernando")

	log.Printf("%+v", e)

	log.Printf("%d", e.getId())

	log.Printf("%s", e.getName())

	log.Println(e.String())
}
