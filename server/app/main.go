package main

import (
	userAmqpInterface "asiap/pkg/user/api/amqp"
	userHttp "asiap/pkg/user/api/http"
	userBusiness "asiap/pkg/user/business"
	userMessagePublisher "asiap/pkg/user/modules/amqp"
	userRepository "asiap/pkg/user/modules/repository"

	notificationAmqpInterface "asiap/pkg/notification/api/amqp"
	notificationBusiness "asiap/pkg/notification/business"
	notificationMessageEmailProvider "asiap/pkg/notification/modules/email"

	amqpCommon "asiap/pkg/common/amqp"

	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var amqpURI = "amqp://guest:guest@rabbitmq:5672/"
var logger = watermill.NewStdLogger(false, false)

func main() {

	amqpConfig := amqp.NewDurableQueueConfig(amqpURI)
	subscriber, err := amqp.NewSubscriber(amqpConfig, logger)
	publisher, err := amqp.NewPublisher(amqpConfig, logger)
	if err != nil {
		panic(err)
	}

	userMsgPub := userMessagePublisher.NewAMQPService(publisher)
	userRepo := userRepository.NewMemoryRepository()
	userService := userBusiness.NewUserService(userRepo, userMsgPub)
	userController := userHttp.NewController(userService)
	userMessageHandlers := userAmqpInterface.NewUserAMQPInterface(userService)

	notificationEmailProvider := notificationMessageEmailProvider.NewEmailMockProvider()
	notificationService := notificationBusiness.NewNotificationService(notificationEmailProvider)
	notificationMessageHandlers := notificationAmqpInterface.NewNotificationAMQPInterface(notificationService)

	/////////////////// Create Message Router
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		// CorrelationID will copy the correlation id from the incoming message's metadata to the produced messages
		middleware.CorrelationID,

		// // The handler function is retried if it returns an error.
		// // After MaxRetries, the message is Nacked and it's up to the PubSub to resend it.
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          logger,
		}.Middleware,

		// // Recoverer handles panics from handlers.
		// // In this case, it passes them as errors to the Retry middleware.
		middleware.Recoverer,
	)
	amqpCommon.RegisterMessageHandler(router, subscriber, publisher, userMessageHandlers)
	amqpCommon.RegisterMessageHandler(router, subscriber, publisher, notificationMessageHandlers)

	router.AddNoPublisherHandler(
		"print_outgoing_messages_in_notification",
		"notification.emailNotificationSent",
		subscriber,
		printMessages,
	)

	ctx := context.Background()
	go router.Run(ctx)
	<-router.Running()
	log.Printf("Message Router is running")

	////////////////// Create HTTP Router

	//create echo http
	e := echo.New()

	//register API path and handler
	userHttp.RegisterPath(e, userController)

	// run server
	go func() {
		// address := fmt.Sprintf("localhost:%d", config.Port)

		if err := e.Start(":3001"); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func printMessages(msg *message.Message) error {
	fmt.Printf(
		"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}
