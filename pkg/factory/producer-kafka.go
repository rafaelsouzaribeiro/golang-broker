package factory

import (
	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/apache-kafka/producer"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

const (
	APACHE_KAFKA_PRODUCER = "apache_kafka_producer"
)

func IProducerBroker(types string, addrs *[]string, message *payload.Message, config *sarama.Config, callback producer.MessageCallback) types.IbrokerKafkaProducer {

	switch types {
	case APACHE_KAFKA_PRODUCER:
		return producer.NewProducer(addrs, message, config, callback)
	}

	return nil
}
