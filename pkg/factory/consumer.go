package factory

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/apache-kafka/consumer"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func IConsumerBroker(broker *[]string, data *payload.Message, channel chan<- payload.Message) types.IbrokerKafkaConsumer {

	return consumer.NewBroker(broker, data, channel)

}
