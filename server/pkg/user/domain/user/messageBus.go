package user

type MessageBus interface {
	UserRegistrationRequested(ID string, email string) error
	UserRegistrationApproved(ID string, email string) error
}
