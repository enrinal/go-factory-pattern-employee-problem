package employee

type Staff struct {
	// TODO: Add the Employee struct
	Employee
}

func NewStaff() *Staff {
	staff := &Staff{
		Employee: Employee{
			Name:   "Staff",
			Salary: 500,
		},
	}
	return staff
}

func (s *Staff) GetBonus() float64 {
	return float64(s.Salary) * 0.1
}
