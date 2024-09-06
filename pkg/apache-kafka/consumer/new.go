package consumer

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

type Broker struct {
	broker  *[]string
	data    *payload.Message
	channel chan<- payload.Message
}

func NewBroker(broker *[]string, data *payload.Message, channel chan<- payload.Message) types.IbrokerKafkaConsumer {
	return &Broker{
		broker:  broker,
		data:    data,
		channel: channel,
	}
}
