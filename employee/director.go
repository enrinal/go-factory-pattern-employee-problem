package employee

type Director struct {
	Employee
}

func NewDirector() *Director {
	return &Director{
		Employee: Employee{
			Name:   "Director",
			Salary: 5000,
			Bonus:  30,
		},
	}
}

func (e *Director) GetBonus() float64 {
	return float64((e.Bonus / 100) * float64(e.Salary))
}
