package main

import (
	"RMQ-Worker/errorHandler"
	"RMQ-Worker/global"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	conn := global.ConnectToRabbitMQ()
	defer conn.Close()

	queues := []string{}
	forever := make(chan bool)

	for _, queue := range queues {
		ch, err := conn.Channel()
		errorHandler.GetError(err, "Unable to open channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			queue, //name
			true,  //durable
			false, //delete when unused
			false, //exclusive
			false, //no wait
			nil,   //arguments
		)
		errorHandler.GetError(err, "unable to declare queue "+queue)

		msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
		errorHandler.GetError(err, "Failed to register Consumer "+q.Name)

		go func() {
			for d := range msgs {
				log.Printf("message: %s from %s", d.Body, d.RoutingKey)
			}
		}()
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
