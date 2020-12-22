package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
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

	err = ch.ExchangeDeclare("logs_direct", "direct", true, false, false, false, nil)
	failOnError(err, "Failed to Declare an Exchange.")

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	failOnError(err, "Failed to declare a queue.")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {
		log.Printf("", q.Name, "logs_direct", s)
		err = ch.QueueBind(q.Name, s, "logs_direct", false, nil)
		failOnError(err, "Failed to bind a queue.")
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer.")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("[x]  %s.", d.Body)
		}
	}()
	log.Printf("[*] Waiting for logs. To exit pree CTRL+C.")

	<-forever

}
