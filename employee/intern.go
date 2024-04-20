package employee

// Intern is a struct for intern
// It embeds Employee
// This means that Intern has all the fields and methods of Employee
type Intern struct {
	Employee
	// TODO: Add the Employee struct
}

// NewIntern creates a new intern
// It returns a pointer to the intern
// Creational method
func NewIntern() *Intern {
	return &Intern{
		Employee: Employee{
			Name:   "Intern",
			Salary: 100,
		},
	}
	// TODO: Create a new intern
	// Set the name to "Intern"
	// Set the salary to 100
}

func (I *Intern) GetBonus() float64 {
	return float64(0)
}
