package employee

import "errors"

// GetEmployeeFactory is a factory for employee
// It returns an employee based on the input
// If the employee is not found, it returns an error
func GetEmployeeFactory(empl string) (EmployeeInterface, error) {
	switch empl {
	case "manager":
		return NewManager(), nil
	case "staff":
		return NewStaff(), nil
	case "intern":
		return NewIntern(), nil
	default:
		return nil, errors.New("employee not found")
	}
}
