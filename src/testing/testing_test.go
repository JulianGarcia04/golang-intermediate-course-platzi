package testing

import "testing"

func TestSum(t *testing.T) {
	//total := Sum(5, 5)
	//
	//if total != 11 {
	//	t.Errorf("Sum was incorrect, got: %d, expected: %d.", total, 11)
	//}

	tables := []struct {
		a, b, n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, table := range tables {
		result := Sum(table.a, table.b)

		if result != table.n {
			t.Errorf("Sum(%d, %d) = %d, want %d", table.a, table.b, result, table.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a, b, n int
	}{
		{2, 1, 2},
		{3, 2, 3},
		{2, 5, 5},
	}
	for _, table := range tables {
		max := GetMax(table.a, table.b)

		if max != table.n {
			t.Errorf("GetMax(%d, %d) = %d, want %d", table.a, table.b, max, table.n)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a, n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, table := range tables {
		r := Fibonacci(table.a)

		if r != table.n {
			t.Errorf("Fibonacci(%d) = %d, want %d", table.a, r, table.n)
		}
	}
}
