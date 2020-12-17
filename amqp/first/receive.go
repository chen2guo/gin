package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, mes string) {
	if err != nil {
		log.Fatalf("%s:   %s", mes, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://cheng:Cheng6688@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel.")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue.")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a aconsumer.")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s.", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C.")
	<-forever

}
