package employee

// Manager adalah struktur untuk karyawan manajer
type Manager struct {
	Name   string
	Salary int
}

// GetName mengembalikan nama karyawan
func (m *Manager) GetName() string {
	return m.Name
}

// GetSalary mengembalikan gaji karyawan
func (m *Manager) GetSalary() int {
	return m.Salary
}

// GetBonus mengembalikan bonus karyawan
func (m *Manager) GetBonus() int {
	return m.Salary * 20 / 100
}

// NewManager adalah factory function untuk membuat karyawan manajer baru
func NewManager() EmployeeInterface {
	return &Manager{
		Name:   "Manager",
		Salary: 5000,
	}
}
