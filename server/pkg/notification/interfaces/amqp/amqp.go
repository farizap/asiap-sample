package amqp

import (
	amqpCommon "asiap/pkg/common/amqp"
	"asiap/pkg/notification/application"
)

func NewNotificationAMQPInterface(s application.NotificationService) []amqpCommon.MessageHandler {
	userCreatedHandler := NewUserCreatedHandler(s)

	return []amqpCommon.MessageHandler{userCreatedHandler}
}
