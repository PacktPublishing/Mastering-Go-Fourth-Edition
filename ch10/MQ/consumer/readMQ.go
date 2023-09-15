package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ consumer")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume("Go", "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received: %s\n", d.Body)
		}
	}()

	fmt.Println("Connected to the RabbitMQ server!")
	<-forever
}
