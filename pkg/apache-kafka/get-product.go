package apachekafka

import (
	"strings"

	"github.com/IBM/sarama"
)

func (p *Broker) GetProducer(config *sarama.Config) (*sarama.AsyncProducer, error) {

	addrs := strings.Split(p.broker, ",")
	producer, err := sarama.NewAsyncProducer(addrs, config)

	if err != nil {
		return nil, err
	}

	return &producer, err
}
