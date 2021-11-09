package amqp

import (
	amqpCommon "asiap/pkg/common/amqp"
	"asiap/pkg/user/business"
)

func NewUserAMQPInterface(s business.UserService) []amqpCommon.MessageHandler {
	userApprovedHandler := NewUserApprovedHandler(s)
	userApprovedTwoHandler := NewUserApprovedTwoHandler(s)

	return []amqpCommon.MessageHandler{userApprovedHandler, userApprovedTwoHandler}
}
