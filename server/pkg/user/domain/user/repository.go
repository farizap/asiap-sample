package user

import "errors"

var ErrNotFound = errors.New("order not found")

type Repository interface {
	Save(*User) error
	ByManagerID(managerId string) (*User, error)
}
