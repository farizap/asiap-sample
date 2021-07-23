package application

import (
	"asiap/pkg/user/domain/user"
	"log"

	"github.com/pkg/errors"
)

type messageService interface {
	UserRegistrationRequested(ID string, email string) error
}

type UserService struct {
	userRegistrationRequestRepository user.Repository
	messagePublisher                  messageService
}

func NewUserService(userRegistrationRequestRepository user.Repository, messageService messageService) UserService {
	return UserService{userRegistrationRequestRepository, messageService}
}

type AddUserRegistrationCommand struct {
	ID        string
	Name      string
	Email     string
	Location  string
	ManagerID string
}

func (s UserService) AddUserRegistration(cmd AddUserRegistrationCommand) error {

	newUserRegistration, err := user.NewUser(cmd.ID, cmd.Name, cmd.Email, cmd.Location, cmd.ManagerID)
	if err != nil {
		return errors.Wrap(err, "cannot create order")
	}

	if err := s.userRegistrationRequestRepository.Save(newUserRegistration); err != nil {
		return errors.Wrap(err, "cannot save order")
	}

	if err := s.messagePublisher.UserRegistrationRequested(newUserRegistration.ID(), newUserRegistration.Email()); err != nil {
		return errors.Wrap(err, "cannot emit event userRegistrationRequested")
	}

	log.Printf("user registration placed", cmd.ID, cmd.Email)

	return nil
}

func (s UserService) UserByManagerID(managerID string) (user.User, error) {
	o, err := s.userRegistrationRequestRepository.ByManagerID(managerID)
	if err != nil {
		return user.User{}, errors.Wrapf(err, "cannot get user %s", managerID)
	}

	return *o, nil
}
