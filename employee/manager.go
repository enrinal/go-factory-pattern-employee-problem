package employee

type Manager struct {
	// TODO: Add the Employee struct
	Employee
}

func NewManager() *Manager {
	manager := &Manager{
		Employee: Employee{
			Name:   "Manager",
			Salary: 1000,
		},
	}
	return manager
}

func (m *Manager) GetBonus() float64 {
	return float64(m.Salary) * 0.2
}
