package factory

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/sns-sqs/sqs"
)

const (
	SQS = "sqs"
)

func ISQSBroker(types string, Configs *payload.SNSSQSMessage, MessageChan chan<- payload.SNSSQSMessage) types.SQS {

	switch types {
	case SQS:
		return sqs.NewBroker(Configs, MessageChan)
	}

	return nil
}
