package factory

import (
	apachekafka "github.com/rafaelsouzaribeiro/golang-broker/pkg/apache-kafka"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/redis"
)

const (
	Kafka = "kafka"
	Redis = "redis"
)

func NewBroker(brokerType string, broker string) types.Ibroker {

	switch brokerType {
	case "kafka":
		return apachekafka.NewBroker(broker)
	case "redis":
		return redis.NewBroker(broker, "123mudar")
	}

	return nil
}
