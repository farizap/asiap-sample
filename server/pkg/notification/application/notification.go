package application

import (
	"asiap/pkg/notification/domain/notification"

	"github.com/pkg/errors"
)

type NotificationService struct {
	notificationProvider notification.Provider
}

func NewNotificationService(p notification.Provider) NotificationService {
	return NotificationService{p}
}

func (s NotificationService) SendEmailNotification(email string) error {

	if err := s.notificationProvider.SendEmail(email); err != nil {
		return errors.Wrap(err, "cannot save order")
	}

	return nil
}
