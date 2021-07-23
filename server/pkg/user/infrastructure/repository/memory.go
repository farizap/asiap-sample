package repository

import "asiap/pkg/user/domain/user"

type MemoryRepository struct {
	users []user.User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{[]user.User{}}
}

func (m *MemoryRepository) Save(orderToSave *user.User) error {
	for i, p := range m.users {
		if p.ID() == orderToSave.ID() {
			m.users[i] = *orderToSave
			return nil
		}
	}

	m.users = append(m.users, *orderToSave)
	return nil
}

func (m MemoryRepository) ByManagerID(managerID string) (*user.User, error) {
	for _, p := range m.users {
		if p.ManagerID() == managerID {
			return &p, nil
		}
	}

	return nil, user.ErrNotFound
}
