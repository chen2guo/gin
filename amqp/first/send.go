package main

import (
	"github.com/streadway/amqp"
	"log"
	"time"
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

	body := "Hello World and World!"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            []byte(body),
		})
	log.Printf("[x] Sent %s.", body)
	failOnError(err, "Failed to publish amessage.")

}
