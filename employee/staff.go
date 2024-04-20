package employee

// Staff is a struct for staff
// It embeds Employee
// This means that Staff has all the fields and methods of Employee
type Staff struct {
	// TODO: Add the Employee struct
	Employee
}

// NewStaff creates a new staff
// It returns a pointer to the staff
// Creational method
func NewStaff() *Staff {
	return &Staff{
		Employee: Employee{
			Name:   "Staff",
			Salary: 500,
			Bonus:  10,
		},
	}
	// TODO: Create a new staff
	// Set the name to "Staff"
	// Set the salary to 500
}

func (e *Staff) GetBonus() float64 {
	return float64((e.Bonus / 100) * float64(e.Salary))
}
