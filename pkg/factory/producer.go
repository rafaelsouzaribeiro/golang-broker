package factory

import (
	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/apache-kafka/producer"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func IProducerBroker(addrs *[]string, message *payload.Message, config *sarama.Config, callback producer.MessageCallback) types.IbrokerKafkaProducer {

	return producer.NewProducer(addrs, message, config, callback)
}
