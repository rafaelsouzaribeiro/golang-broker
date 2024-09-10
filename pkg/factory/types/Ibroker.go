package types

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

type Ibroker interface {
	Consumer(data *payload.Message, channel chan<- payload.Message)
	ListenPartition(data *payload.Message, channel chan<- payload.Message)
	SendMessage(data *payload.Message)
}
