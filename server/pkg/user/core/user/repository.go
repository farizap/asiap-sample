package user

import "errors"

var ErrNotFound = errors.New("user not found")

type Repository interface {
	Save(*User) error
	ByManagerID(managerId string) (*[]User, error)
	ByID(id string) (*User, error)
}
