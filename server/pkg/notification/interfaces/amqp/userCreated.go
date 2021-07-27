package amqp

import (
	"asiap/pkg/common/event"
	"asiap/pkg/notification/application"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
)

type UserCreatedHandler struct {
	service application.NotificationService
}

func NewUserCreatedHandler(s application.NotificationService) *UserCreatedHandler {
	return &UserCreatedHandler{s}
}

func (h UserCreatedHandler) Name() string {
	return "notification_userCreatedHandler"
}

func (h UserCreatedHandler) SubscribeTopic() string {
	return event.UserCreated
}

func (h UserCreatedHandler) PublishTopic() string {
	return event.EmailNotificationSent
}

func (h UserCreatedHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	p := event.UserRegistrationCreatedMsg{}
	err := json.Unmarshal(msg.Payload, &p)
	if err != nil {
		log.Fatal(err)
	}

	h.service.SendEmailNotification(p.Email)

	userRegistration := event.EmailNotificationSentMsg{
		ID:    p.ID,
		Email: p.Email,
	}

	b, err := json.Marshal(userRegistration)
	if err != nil {
		return nil, errors.Wrap(err, "cannot marshal userRegistration for amqp")
	}

	msg = message.NewMessage(watermill.NewUUID(), b)

	return message.Messages{msg}, nil
}
