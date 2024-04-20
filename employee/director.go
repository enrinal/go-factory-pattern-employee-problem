package employee

// Manager is a struct for manager
// It embeds Employee
// This means that Manager has all the fields and methods of Employee
type Director struct {
	// TODO: Add the Employee struct
	Employee
}

// NewManager creates a new manager
// It returns a pointer to the manager
// Creational method
func NewDirector() *Director {
	return &Director{
		Employee: Employee{
			Name:   "Director",
			Salary: 5000,
		},
	}
}

func (d *Director) GetBonus() float64 {
	return float64(d.Salary) * 0.3
}
