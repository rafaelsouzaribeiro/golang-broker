package types

import (
	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

type Ibroker interface {
	Consumer(data *payload.Message, channel chan<- payload.Message)
	ListenPartition(data *payload.Message, channel chan<- payload.Message)
	GetProducer(config *sarama.Config) (*sarama.AsyncProducer, error)
	SendMessage(producer *sarama.AsyncProducer, data *payload.Message)
}
