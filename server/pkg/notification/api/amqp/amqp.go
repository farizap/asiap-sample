package amqp

import (
	amqpCommon "asiap/pkg/common/amqp"
	"asiap/pkg/notification/business"
)

func NewNotificationAMQPInterface(s business.NotificationService) []amqpCommon.MessageHandler {
	userCreatedHandler := NewUserCreatedHandler(s)

	return []amqpCommon.MessageHandler{userCreatedHandler}
}
