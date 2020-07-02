package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	go server()
	go client()

	var s string
	fmt.Scanln(&s)
}

func client() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Error receiving from queue")
	for msg := range msgs {
		fmt.Printf("Message received: %s\n", msg.Body)
	}
}

func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello RabbitMQ"),
	}
	for {

		ch.Publish("", q.Name, false, false, msg)
	}
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Unable to dial connection")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Cannot create a queue")
	return conn, ch, &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		log.Panic(err)
	}
}
