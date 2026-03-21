package main

import (
	"fmt"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func main() {
	godotenv.Load()
	broker := factory.IRabbitMQBroker()

	channel, err := broker.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer channel.Close()

	ch := make(chan *amqp.Delivery)
	payload := payload.RabbitMQMessage{
		QueueName: "test-queue",
		Channel:   channel,
		Msgs:      ch,
	}

	go broker.Consume(&payload)

	for msg := range payload.Msgs {
		fmt.Println(string(msg.Body))

		if value, ok := msg.Headers["header"].(string); ok {
			fmt.Printf("Header:%s", value)
		}

		msg.Ack(false)
	}

	close(payload.Msgs)

}
