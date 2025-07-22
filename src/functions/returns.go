package functions

import "fmt"

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func Main() {
	total1 := sum(1, 2, 3)
	total2 := sum(1, 2, 3, 4)
	total3 := sum(1, 2, 3, 4, 5)

	fmt.Println(total1, total2, total3)

	fmt.Println(getValues(2))
}
