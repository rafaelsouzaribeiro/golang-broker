package sns

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

type Broker struct {
	Config *payload.SNSSQSMessage
}

func NewBroker(config *payload.SNSSQSMessage) types.SNS {
	return &Broker{
		Config: config,
	}
}
