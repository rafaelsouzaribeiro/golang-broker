package producer

import (
	"github.com/IBM/sarama"
)

type Producer struct {
	addrs   []string
	topic   string
	message sarama.Encoder
	config  *sarama.Config
}

func NewProducer(addrs []string, topic string, message sarama.Encoder, config *sarama.Config) *Producer {
	return &Producer{
		addrs:   addrs,
		topic:   topic,
		message: message,
		config:  config,
	}
}

func (p *Producer) GetProducer() (*sarama.AsyncProducer, error) {

	producer, err := sarama.NewAsyncProducer(p.addrs, p.config)

	if err != nil {
		return nil, err
	}

	return &producer, err
}

func (p *Producer) SendMessage(producer *sarama.AsyncProducer) error {
	message := &sarama.ProducerMessage{Topic: p.topic, Value: p.message}

	(*producer).Input() <- message

	err := <-(*producer).Errors()

	if err != nil {
		panic(err)
	}

	return nil
}
