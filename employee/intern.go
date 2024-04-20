package employee

// Intern is a struct for intern
// It embeds Employee
// This means that Intern has all the fields and methods of Employee
type Intern struct {
	// TODO: Add the Employee struct
	Employee
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
}
func (i *Intern) GetBonus() float64 {
	return float64(i.Salary) * 0.0
}
