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
		},
	}
}

func (d *Director) GetBonus() float64 {
	return float64(d.Salary) * 0.3
}
