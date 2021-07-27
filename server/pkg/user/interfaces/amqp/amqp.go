package amqp

import (
	amqpCommon "asiap/pkg/common/amqp"
	"asiap/pkg/user/application"
)

func NewUserAMQPInterface(s application.UserService) []amqpCommon.MessageHandler {
	userApprovedHandler := NewUserApprovedHandler(s)

	return []amqpCommon.MessageHandler{userApprovedHandler}
}
