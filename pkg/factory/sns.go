package factory

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/sns-sqs/sns"
)

func ISNSBroker(config *payload.SNSSQSMessage) types.SNS {

	return sns.NewBroker(config)

}
