package producer

import (
	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

type MessageCallback func(messages payload.Message)

type Producer struct {
	addrs    []string
	message  payload.Message
	config   *sarama.Config
	callback MessageCallback
}

func NewProducer(addrs *[]string, message *payload.Message, config *sarama.Config, callback MessageCallback) types.IbrokerKafkaProducer {
	return &Producer{
		addrs:    *addrs,
		message:  *message,
		config:   config,
		callback: callback,
	}
}
