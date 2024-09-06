package sqs

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

type Broker struct {
	Configs     *payload.SNSSQSMessage
	MessageChan chan<- payload.SNSSQSMessage
}

func NewBroker(Configs *payload.SNSSQSMessage, MessageChan chan<- payload.SNSSQSMessage) types.SQS {
	return &Broker{
		Configs:     Configs,
		MessageChan: MessageChan,
	}
}
