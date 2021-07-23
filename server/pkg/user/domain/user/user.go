package user

import "errors"

var ErrEmptyOrderID = errors.New("empty user id")

type User struct {
	id        string
	name      string
	email     string
	location  string
	managerID string
}

func (o User) ID() string {
	return o.id
}

func (o User) Name() string {
	return o.Name()
}

func (o User) Email() string {
	return o.email
}

func (o User) Location() string {
	return o.location
}

func (o User) ManagerID() string {
	return o.managerID
}

func NewUser(id string, name string, email string, location string, managerID string) (*User, error) {
	if len(id) == 0 {
		return nil, ErrEmptyOrderID
	}

	return &User{id, name, email, location, managerID}, nil
}
