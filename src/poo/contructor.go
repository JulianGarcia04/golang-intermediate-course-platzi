package poo

import "log"

func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}

func Constructor() {
	// contructor

	// 1
	e1 := Employee{}

	log.Println(e1.String())

	// 2
	e2 := Employee{
		id:   1,
		name: "Jack",
	}

	log.Println(e2.String())

	// 3
	e3 := new(Employee)

	log.Println(e3.String())

	//4 RECOMMENDED
	e4 := NewEmployee(4, "Bob")

	log.Println(e4.String())
}
