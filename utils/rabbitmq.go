package utils

import (
	"log"

	"github.com/streadway/amqp"
)

var RabbitMq *amqp.Connection

func ConnectMessageBroker(addr string) {
	conn, err := amqp.Dial(addr)

	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}
	RabbitMq = conn
}
