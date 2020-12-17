package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
	"time"
)

func bodyFrom(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "Hello  World"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

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

	if err != nil {
		fmt.Printf("Err Queue failed, err: %v\n", err)
		return
	}
	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    2,
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
	if err != nil {
		fmt.Printf("Publish a message Failed , err:%v\n.", err)
		return
	}
	log.Printf("[X] Sent  %s.", body)
}
