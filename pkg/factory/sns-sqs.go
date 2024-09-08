package factory

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
	snssqs "github.com/rafaelsouzaribeiro/golang-broker/pkg/sns-sqs"
)

func ISQSSNSBroker(config *payload.SNSSQSMessage) types.SQSSNS {

	return snssqs.NewBroker(config)

}
