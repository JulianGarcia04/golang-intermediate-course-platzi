package mock_testing

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Dni:  "1",
						Name: "Pedro",
						Age:  20,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Dni:  "1",
					Name: "Pedro",
					Age:  20,
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, tt := range table {
		tt.mockFunc()

		ft, err := GetFullTimeEmployeeById(tt.id, tt.dni)

		if err != nil {
			t.Errorf("GetFullTimeEmployeeById(%d, %s): error %v", tt.id, tt.dni, err)
		}

		if ft.Age != tt.expectedEmployee.Age {
			t.Errorf("Expected: %d, got: %d", tt.expectedEmployee.Age, ft.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
