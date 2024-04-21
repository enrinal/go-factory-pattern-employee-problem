package employee

// EmployeeInterface adalah interface untuk objek Employee
type EmployeeInterface interface {
	GetName() string
	GetSalary() int
	GetBonus() int
}

// Factory adalah function yang mengembalikan instance dari tipe EmployeeInterface
type Factory func() EmployeeInterface
