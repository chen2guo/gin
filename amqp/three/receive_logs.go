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

	err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	failOnError(err, "Failed to Declare an Exchange.")

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	failOnError(err, "Failed to declare a queue.")

	err = ch.QueueBind(q.Name, "", "logs", false, nil)
	failOnError(err, "Failed to bind a queue.")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer.")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] receive %s.", d.Body)
		}
	}()
	log.Printf("[*] Waiting for logs. To exit press CTRL+C.")

	<-forever

}
