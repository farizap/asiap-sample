package main

import (
	userApplication "asiap/pkg/user/application"
	userMessagePublisher "asiap/pkg/user/infrastructure/amqp"
	userRepository "asiap/pkg/user/infrastructure/repository"
	userHttp "asiap/pkg/user/interfaces/http"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var amqpURI = "amqp://guest:guest@rabbitmq:5672/"

func main() {

	amqpConfig := amqp.NewDurableQueueConfig(amqpURI)
	subscriber, err := amqp.NewSubscriber(
		// This config is based on this example: https://www.rabbitmq.com/tutorials/tutorial-two-go.html
		// It works as a simple queue.
		//
		// If you want to implement a Pub/Sub style service instead, check
		// https://watermill.io/pubsubs/amqp/#amqp-consumer-groups
		amqpConfig,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	///////
	go process(messages)
	///////

	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		panic(err)
	}

	userMsgPub := userMessagePublisher.NewAMQPService(publisher)
	userRepo := userRepository.NewMemoryRepository()
	userService := userApplication.NewUserService(userRepo, userMsgPub)
	userController := userHttp.NewController(userService)

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

	//close db
	// defer dbCon.CloseConnection()

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
