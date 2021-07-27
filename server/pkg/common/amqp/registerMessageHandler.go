package amqp

import (
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

type MessageHandler interface {
	Handler(msg *message.Message) ([]*message.Message, error)
	Name() string
	SubscribeTopic() string
	PublishTopic() string
}

func RegisterMessageHandler(router *message.Router, subscriber *amqp.Subscriber, publisher *amqp.Publisher, handlers []MessageHandler) {
	for _, handler := range handlers {
		router.AddHandler(
			handler.Name(),           // handler name, must be unique
			handler.SubscribeTopic(), // topic from which we will read events
			subscriber,
			handler.PublishTopic(), // topic to which we will publish events
			publisher,
			handler.Handler,
		)
	}
}
