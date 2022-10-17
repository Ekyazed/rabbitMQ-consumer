package global

import (
	"RMQ-Worker/errorHandler"

	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitMQ() *amqp.Connection {

	username := os.Getenv("AMQP_USERNAME")
	password := os.Getenv("AMQP_PASSWORD")
	host := os.Getenv("AMQP_HOST")
	port := os.Getenv("AMQP_PORT")

	conn, err := amqp.Dial("amqp://" + username + ":" + password + "@" + host + ":" + port + "/")
	errorHandler.GetError(err, "Failed to open dial")

	return conn
}
