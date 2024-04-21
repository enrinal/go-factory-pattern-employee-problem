package employee

type Intern struct {
	// TODO: Add the Employee struct
	Employee
}

func NewIntern() *Intern {
	intern := &Intern{
		Employee: Employee{
			Name:   "Intern",
			Salary: 100,
		},
	}
	return intern
}

func (i *Intern) GetBonus() float64 {
	return float64(i.Salary) * 0.0
}
