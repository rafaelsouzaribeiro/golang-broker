package types

import (
	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

type IbrokerKafkaProducer interface {
	GetProducer() (*sarama.AsyncProducer, error)
	SendMessage(producer *sarama.AsyncProducer)
	GetErrorMessage() *payload.Message
}
