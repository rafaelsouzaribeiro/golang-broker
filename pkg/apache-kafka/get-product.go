package apachekafka

import (
	"strings"

	"github.com/IBM/sarama"
)

func GetProducer(config *sarama.Config, broker string) (*sarama.AsyncProducer, error) {

	addrs := strings.Split(broker, ",")
	producer, err := sarama.NewAsyncProducer(addrs, config)

	if err != nil {
		return nil, err
	}

	return &producer, err
}
