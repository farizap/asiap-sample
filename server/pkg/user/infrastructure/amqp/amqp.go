package amqp

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
)

type UserRegistrationRequestedMsg struct {
	ID        string `json:"id"`
	ManagerID string `json:"managerID"`
}

type AMQPPublisher struct {
	publisher message.Publisher
}

func NewAMQPService(publisher message.Publisher) AMQPPublisher {
	return AMQPPublisher{publisher}
}

func (i AMQPPublisher) UserRegistrationRequested(id string, managerID string) error {
	userRegistration := UserRegistrationRequestedMsg{
		ID:        id,
		ManagerID: managerID,
	}

	b, err := json.Marshal(userRegistration)
	if err != nil {
		return errors.Wrap(err, "cannot marshal userRegistration for amqp")
	}

	msg := message.NewMessage(watermill.NewUUID(), b)

	err = i.publisher.Publish(
		"example.topic",
		msg,
	)
	if err != nil {
		return errors.Wrap(err, "cannot send userRegistrationRequested to amqp")
	}

	log.Printf("sent order %s to amqp", id)

	return nil
}
