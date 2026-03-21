package factory

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/rabbitmq"
)

func IRabbitMQBroker() types.IbrokerRabbitMQ {
	return rabbitmq.NewRabbitMQBroker()
}
