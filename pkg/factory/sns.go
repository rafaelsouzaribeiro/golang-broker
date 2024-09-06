package factory

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/sns-sqs/sns"
)

const (
	SNS = "sns"
)

func ISNSBroker(types string, config *payload.SNSSQSMessage) types.SNS {

	switch types {
	case SNS:
		return sns.NewBroker(config)
	}

	return nil
}
