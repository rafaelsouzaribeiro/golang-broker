package factory

import (
	apachekafka "github.com/rafaelsouzaribeiro/golang-broker/pkg/apache-kafka"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
)

const (
	Kafka = "kafka"
)

func NewBroker(brokerType string, broker string) types.Ibroker {

	switch brokerType {
	case "kafka":
		return apachekafka.NewBroker(broker)
	}

	return nil
}
