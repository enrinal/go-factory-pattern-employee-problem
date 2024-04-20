package employee

import "errors"

// GetEmployeeFactory is a factory for employee
// It returns an employee based on the input
// If the employee is not found, it returns an error
func GetEmployeeFactory(empl string) (EmployeeInterface, error) {

	if empl == "manager" {
		return NewManager(), nil
		// TODO: Return a new manager
	}

	if empl == "staff" {
		return NewStaff(), nil
		// TODO: Return a new staff
	}

	if empl == "intern" {
		return NewIntern(), nil
		// TODO: Return a new intern
	}

	// TODO: Create director object

	return nil, errors.New("Employee not found")
}
