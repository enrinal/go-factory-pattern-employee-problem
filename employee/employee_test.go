package employee_test

import (
	"go-factorypattern-company-case/employee"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employee", func() {

	Context("Manager Object", func() {

		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("manager")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Manager"))
			Expect(empl.GetSalary()).To(Equal(1000))
		})

		It("should return the correct bonus", func() {
			manager := employee.NewManager()
			// Calculate expected bonus
			expectedBonus := float64(manager.GetSalary()) * 0.2

			Expect(manager.GetBonus()).To(Equal(expectedBonus))
		})

	})

	Context("Staff Object", func() {

		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("staff")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Staff"))
			Expect(empl.GetSalary()).To(Equal(500))
		})

		It("should return the correct bonus", func() {
			// TODO Implement the test for the bonus
			// Salary is 500
			// Bonus is 10% of the salary
			// Bonus is 50
			staff := employee.NewStaff()
			// Calculate expected bonus
			expectedBonus := float64(staff.GetSalary()) * 0.1

			Expect(staff.GetBonus()).To(Equal(expectedBonus))

		})

	})

	Context("Intern Object", func() {

		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("intern")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Intern"))
			Expect(empl.GetSalary()).To(Equal(100))
		})

		It("should return the correct bonus", func() {
			intern := employee.NewIntern()
			// Calculate expected bonus
			expectedBonus := float64(intern.GetSalary()) * 0.0

			Expect(intern.GetBonus()).To(Equal(expectedBonus))
		})
	})

	// TODO: Implement the test for the Director object
	Context("Director Object", func() {
		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("director")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Director"))
			Expect(empl.GetSalary()).To(Equal(5000))
		})

		It("should return the correct bonus", func() {
			director := employee.NewDierector()
			// Calculate expected bonus
			expectedBonus := float64(director.GetSalary()) * 0.3

			Expect(director.GetBonus()).To(Equal(expectedBonus))
		})
	})

	Context("Empty Employee", func() {

		It("should return an error", func() {
			_, err := employee.GetEmployeeFactory("non-existing")
			Expect(err).To(HaveOccurred())
		})

	})

})
