package main

import (
	"fmt"
	"go-factorypattern-company-case/employee"
)

// use main function to demonstrate the factory pattern
func main() {

	budi, err := employee.GetEmployeeFactory("manager")
	if err != nil {
		panic(err)
	}

	printDetails(budi)

}

func printDetails(e employee.EmployeeInterface) {
	fmt.Printf("Employee: %s", e.GetName())
	fmt.Println()
	fmt.Printf("Salary: %d", e.GetSalary())
	fmt.Println()
	fmt.Printf("Bonus: %.2f", e.GetBonus())
	fmt.Println()
}
