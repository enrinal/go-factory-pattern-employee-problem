package employee

type Director struct {
	Employee
}

func NewDirector() *Director {
	return &Director{
		Employee: Employee{
			Name:   "Director",
			Salary: 5000,
		},
	}
}

func (d *Director) GetBonus() float64 {
	return float64(d.Salary * 30 / 100)
}
