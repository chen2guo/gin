package main

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://cheng:Cheng6688@localhost:5672/")
	if err != nil {
		fmt.Printf("Connect to RabbitMQ Failed ,err: %v\n.", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Open a channel failed, err:%v\n.", err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil)

	err = ch.Qos(1, 0, false)
	if err != nil {
		fmt.Printf("ch.Qos() failed, err: %v\n.", err)
		return
	}

	megs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Printf("ch.Consume failed, err; %v\n.", err)
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range megs {
			log.Printf("Received  a message: %s.", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done.")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C.")
	<-forever
}
