package factory

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/apache-kafka/consumer"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

const (
	APACHE_KAFKA_CONSUMER = "apache_kafka_consumer"
	SNS_SQS               = "sns_sqs"
)

func INewBroker(types string, broker *[]string, data *payload.Message, channel chan<- payload.Message) types.IbrokerKafkaConsumer {

	switch types {
	case APACHE_KAFKA_CONSUMER:
		return consumer.NewBroker(broker, data, channel)
	}

	return nil
}
