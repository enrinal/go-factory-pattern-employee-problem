package employee

// Intern adalah struktur untuk karyawan intern
type Intern struct {
	Name   string
	Salary int
}

// GetName mengembalikan nama karyawan
func (i *Intern) GetName() string {
	return i.Name
}

// GetSalary mengembalikan gaji karyawan
func (i *Intern) GetSalary() int {
	return i.Salary
}

// GetBonus mengembalikan bonus karyawan
func (i *Intern) GetBonus() int {
	return 0
}

// NewIntern adalah factory function untuk membuat karyawan intern baru
func NewIntern() EmployeeInterface {
	return &Intern{
		Name:   "Intern",
		Salary: 2000,
	}
}
