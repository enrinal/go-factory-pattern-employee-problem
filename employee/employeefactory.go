package employee

import "errors"

func GetEmployeeFactory(empl string) (EmployeeInterface, error) {

	if empl == "manager" {
		// TODO: Return a new manager
		return NewManager(), nil
	}

	if empl == "staff" {
		// TODO: Return a new staff
		return NewStaff(), nil
	}

	if empl == "intern" {
		// TODO: Return a new intern
		return NewIntern(), nil
	}

	// TODO: Create director object
	if empl == "director" {
		return NewDirector(), nil
	}

	return nil, errors.New("Employee not found")
}
