package types

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

type IbrokerRabbitMQ interface {
	Publish(payload *payload.RabbitMQMessage) error
	Consume(payload *payload.RabbitMQMessage) error
	OpenChannel() (*amqp.Channel, error)
}
