package amqp

import (
	"asiap/pkg/common/event"
	"asiap/pkg/user/business"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
)

type UserApprovedTwoHandler struct {
	service business.UserService
}

func NewUserApprovedTwoHandler(s business.UserService) *UserApprovedTwoHandler {
	return &UserApprovedTwoHandler{s}
}

func (h UserApprovedTwoHandler) Name() string {
	return "user_userApprovedTwoHandlerTwo"
}

func (h UserApprovedTwoHandler) SubscribeTopic() string {
	return event.UserCreated
}

func (h UserApprovedTwoHandler) PublishTopic() string {
	return event.EmailNotificationSent
}

func (h UserApprovedTwoHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	p := event.UserRegistrationApprovedMsg{}
	err := json.Unmarshal(msg.Payload, &p)
	if err != nil {
		log.Fatal(err)
	}

	h.service.ActivateUser(p.ID)

	userRegistration := event.UserRegistrationCreatedMsg{
		ID:    p.ID,
		Email: p.Email,
	}

	b, err := json.Marshal(userRegistration)
	if err != nil {
		return nil, errors.Wrap(err, "cannot marshal userRegistration for amqp")
	}

	log.Printf("sent userCreatedTwo event with user id: %s to amqp", p.ID)

	msg = message.NewMessage(watermill.NewUUID(), b)

	return message.Messages{msg}, nil
}
