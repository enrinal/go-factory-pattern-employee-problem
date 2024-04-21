package employee

type Director struct {
	// TODO: Add the Employee struct
	Employee
}

func NewDirector() *Director {
	// TODO: Create a new director
	// Set the name to "Director"
	// Set the salary to 5000
	director := &Director{
		Employee: Employee{
			Name:   "Director",
			Salary: 5000,
			Bonus:  5000.0 * 0.3,
		},
	}
	return director
}
