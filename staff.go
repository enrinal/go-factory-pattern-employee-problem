package employee

// Staff adalah struktur untuk karyawan staf
type Staff struct {
	Name   string
	Salary int
}

// GetName mengembalikan nama karyawan
func (s *Staff) GetName() string {
	return s.Name
}

// GetSalary mengembalikan gaji karyawan
func (s *Staff) GetSalary() int {
	return s.Salary
}

// GetBonus mengembalikan bonus karyawan
func (s *Staff) GetBonus() int {
	return s.Salary * 10 / 100
}

// NewStaff adalah factory function untuk membuat karyawan staf baru
func NewStaff() EmployeeInterface {
	return &Staff{
		Name:   "Staff",
		Salary: 3000,
	}
}
