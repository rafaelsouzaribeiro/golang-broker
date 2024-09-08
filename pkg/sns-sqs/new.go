package snssqs

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

type Broker struct {
	Config *payload.SNSSQSMessage
}

func NewBroker(config *payload.SNSSQSMessage) *Broker {
	return &Broker{
		Config: config,
	}
}
