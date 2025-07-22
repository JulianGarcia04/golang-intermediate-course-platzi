package poo

import "fmt"

type TemporaryEmployee struct {
	taxRate string
}

func (t TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	getMessage() string
}

func PrintMessage(e PrintInfo) {
	fmt.Println(e.getMessage())
}

func Interfaces() {
	te := new(TemporaryEmployee)

	PrintMessage(te)
}
