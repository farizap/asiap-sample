package business

import (
	"asiap/pkg/user/core/user"

	"github.com/pkg/errors"
)

type UserService struct {
	userRegistrationRequestRepository user.Repository
	messagePublisher                  user.MessageBus
}

func NewUserService(userRegistrationRequestRepository user.Repository, messageService user.MessageBus) UserService {
	return UserService{userRegistrationRequestRepository, messageService}
}

type AddUserRegistrationCommand struct {
	ID        string
	Name      string
	Email     string
	Location  string
	ManagerID string
	Status    string
}

func (s UserService) AddUserRegistration(cmd AddUserRegistrationCommand) error {

	newUserRegistration, err := user.NewUser(cmd.ID, cmd.Name, cmd.Email, cmd.Location, cmd.ManagerID)
	if err != nil {
		return errors.Wrap(err, "cannot create order")
	}

	existingUser, _ := s.userRegistrationRequestRepository.ByID(newUserRegistration.ID())
	if existingUser != nil {
		return errors.New("ID already used")
	}

	if err := s.userRegistrationRequestRepository.Save(newUserRegistration); err != nil {
		return errors.Wrap(err, "cannot save order")
	}

	if err := s.messagePublisher.UserRegistrationRequested(newUserRegistration.ID(), newUserRegistration.Email()); err != nil {
		return errors.Wrap(err, "cannot emit event userRegistrationRequested")
	}

	return nil
}

func (s UserService) ApproveRegistration(id string) error {

	existingUser, err := s.userRegistrationRequestRepository.ByID(id)
	if err != nil {
		return errors.Wrap(err, "user not exist")
	}
	existingUser.ApproveUser()

	if err := s.userRegistrationRequestRepository.Save(existingUser); err != nil {
		return errors.Wrap(err, "cannot save order")
	}

	if err := s.messagePublisher.UserRegistrationApproved(existingUser.ID(), existingUser.Email()); err != nil {
		return errors.Wrap(err, "cannot emit event userRegistrationApproved")
	}

	return nil
}

func (s UserService) ActivateUser(id string) error {

	existingUser, err := s.userRegistrationRequestRepository.ByID(id)
	if err != nil {
		return errors.Wrap(err, "user not exist")
	}

	existingUser.ActivateUser()

	if err := s.userRegistrationRequestRepository.Save(existingUser); err != nil {
		return errors.Wrap(err, "cannot save order")
	}

	return nil
}

func (s UserService) UserByManagerID(managerID string) ([]user.User, error) {
	o, err := s.userRegistrationRequestRepository.ByManagerID(managerID)
	if err != nil {
		return []user.User{}, errors.Wrapf(err, "cannot get user %s", managerID)
	}

	return *o, nil
}
