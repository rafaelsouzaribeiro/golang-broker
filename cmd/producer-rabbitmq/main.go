package main

import (
	"github.com/joho/godotenv"
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
	payload := payload.RabbitMQMessage{
		Channel:    channel,
		Exchange:   "amq.direct",
		RoutingKey: "test-routing-key",
		Value:      []byte("Test message"),
		Header:     "Header value",
	}

	broker.Publish(&payload)

}
