package amqp

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/pkg/errors"

	"asiap/pkg/common/event"
)

type AMQPPublisher struct {
	publisher message.Publisher
}

func NewAMQPService(publisher message.Publisher) AMQPPublisher {
	return AMQPPublisher{publisher}
}

func (i AMQPPublisher) UserRegistrationRequested(id string, email string) error {
	userRegistration := event.UserRegistrationRequestedMsg{
		ID:    id,
		Email: email,
	}

	b, err := json.Marshal(userRegistration)
	if err != nil {
		return errors.Wrap(err, "cannot marshal userRegistration for amqp")
	}

	msg := message.NewMessage(watermill.NewUUID(), b)

	err = i.publisher.Publish(
		event.UserRegistrationRequested,
		msg,
	)
	if err != nil {
		return errors.Wrap(err, "cannot send userRegistrationRequested to amqp")
	}

	log.Printf("sent order %s to amqp", id)

	return nil
}

func (i AMQPPublisher) UserRegistrationApproved(id string, email string) error {
	userRegistration := event.UserRegistrationApprovedMsg{
		ID:    id,
		Email: email,
	}

	b, err := json.Marshal(userRegistration)
	if err != nil {
		return errors.Wrap(err, "cannot marshal userRegistration for amqp")
	}

	msg := message.NewMessage(watermill.NewUUID(), b)
	middleware.SetCorrelationID(watermill.NewUUID(), msg)

	err = i.publisher.Publish(
		event.UserRegistrationApproved,
		msg,
	)
	if err != nil {
		return errors.Wrap(err, "cannot send userRegistrationRequested to amqp")
	}

	log.Printf("sent userApproved event with user id: %s to amqp", id)

	return nil
}
