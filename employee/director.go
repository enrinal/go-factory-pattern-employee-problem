package employee

type Director struct {
	// TODO: Add the Employee struct
	Employee
}

func NewDierector() *Director {

	return &Director{
		Employee: Employee{
			Name:   "Director",
			Salary: 5000,
			Bonus:  0.3 * 5000,
		},
	}
}
