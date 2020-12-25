package main

import (
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:  %s ", msg, err)
	}
}

func readomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	n := bodyFrom(os.Args)

	log.Printf("[x Requesting fib(%d)", n)

	res, err := fibonacciRPC(n)
	failOnError(err, "Failed to hanfle RPC request.")

	log.Printf("[.] Got %d.", res)
}

func fibonacciRPC(n int) (res int, err error) {
	conn, err := amqp.Dial("amqp://cheng:Cheng6688@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // type
		false, // durable
		false, // auto-deleted
		false, // internal
		nil,   // no-wait
	)
	failOnError(err, "Failed to declare an queue")

	err = ch.Qos(1, 0, false)
	failOnError(err, "Failed to set Qos.")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer.")

	corrID := readomString(32)

	err = ch.Publish(
		"", "rpc_queue", false, false,
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         q.Name,
			Expiration:      "",
			MessageId:       corrID,
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            []byte(strconv.Itoa(n)),
		})
	failOnError(err, "Failed to publish a message.")

	for d := range msgs {
		if corrID == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to  convert body to integer.")
			break
		}
	}
	return
}

func bodyFrom(args []string) int {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], "")
	}
	n, err := strconv.Atoi(s)
	failOnError(err, "Failed to convert arg to integger.")
	return n

}
