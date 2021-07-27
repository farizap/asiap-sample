package user

import "errors"

var ErrEmptyUserID = errors.New("empty user id")

type UserStatus string

const (
	Requested UserStatus = "requested"
	Approved             = "approved"
	Rejected             = "rejected"
	Active               = "active"
)

type User struct {
	id        string
	name      string
	email     string
	location  string
	managerID string
	status    UserStatus
}

func (o User) ID() string {
	return o.id
}

func (o User) Name() string {
	return o.name
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

func (o User) Status() string {
	return string(o.status)
}

func (o *User) ApproveUser() {
	o.status = Approved
}

func (o *User) ActivateUser() {
	o.status = Active
}

func NewUser(id string, name string, email string, location string, managerID string) (*User, error) {
	if len(id) == 0 {
		return nil, ErrEmptyUserID
	}

	return &User{id, name, email, location, managerID, Requested}, nil
}
