package main

import (
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func failOnError(err error, mes string) {
	if err != nil {
		log.Fatalf("%s:   %s", mes, err)
	}
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz~!@#$%^&*()"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func bodyFrom(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = GetRandomString(26)
	} else {
		s = strings.Join(args[1:], "")
	}
	return s
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

	body := bodyFrom(os.Args)
	err = ch.Publish("logs", "", false, false, amqp.Publishing{
		Headers:         nil,
		ContentType:     "",
		ContentEncoding: "",
		DeliveryMode:    0,
		Priority:        0,
		CorrelationId:   "text/plain",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "",
		Timestamp:       time.Time{},
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            []byte(body),
	})
	failOnError(err, "Failed to publish a message.")

	time.Sleep(time.Second * 1)

	log.Printf("  [x]  Sent %s.", body)

}
