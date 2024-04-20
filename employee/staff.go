package employee

// Staff is a struct for staff
// It embeds Employee
// This means that Staff has all the fields and methods of Employee
type Staff struct {
	Employee
	// TODO: Add the Employee struct
}

// NewStaff creates a new staff
// It returns a pointer to the staff
// Creational method
func NewStaff() *Staff {
	// TODO: Create a new staff
	// Set the name to "Staff"
	// Set the salary to 500
	return &Staff{
		Employee: Employee{
			Name:   "Staff",
			Salary: 500,
			Bonus:  500 * 0.1,
		},
	}
}
