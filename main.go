package main

import (
	"fmt"
	"go-factorypattern-company-case/employee"
)

// use main function to demonstrate the factory pattern
func main() {
	// Create Manager
	budi, err := employee.GetEmployeeFactory("manager")
	if err != nil {
		panic(err)
	}
	printDetails(budi)

	// Create Staff
	joko, err := employee.GetEmployeeFactory("staff")
	if err != nil {
		panic(err)
	}
	printDetails(joko)

	// Create Intern
	susi, err := employee.GetEmployeeFactory("intern")
	if err != nil {
		panic(err)
	}
	printDetails(susi)

	// Create Director
	dian, err := employee.GetEmployeeFactory("director")
	if err != nil {
		panic(err)
	}
	printDetails(dian)
}

func printDetails(e employee.EmployeeInterface) {
	fmt.Printf("Employee: %s", e.GetName())
	fmt.Println()
	fmt.Printf("Salary: %d", e.GetSalary())
	fmt.Println()
	fmt.Printf("Bonus: %d", e.GetBonus())
	fmt.Println()
}
