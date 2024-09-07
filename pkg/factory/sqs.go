package factory

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/sns-sqs/sqs"
)

func ISQSBroker(Configs *payload.SNSSQSMessage, MessageChan chan<- payload.SNSSQSMessage) types.SQS {

	return sqs.NewBroker(Configs, MessageChan)

}
