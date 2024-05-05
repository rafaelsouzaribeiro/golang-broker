package producer

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/utils"
)

type MessageCallback func(messages string)

type Producer struct {
	addrs    []string
	message  utils.Message
	config   *sarama.Config
	callback MessageCallback
}

func NewProducer(addrs []string, message utils.Message, config *sarama.Config, callback MessageCallback) *Producer {
	return &Producer{
		addrs:    addrs,
		message:  message,
		config:   config,
		callback: callback,
	}
}

func (p *Producer) GetProducer() (*sarama.AsyncProducer, error) {

	producer, err := sarama.NewAsyncProducer(p.addrs, p.config)

	if err != nil {
		p.callback(p.GetErrorMessage())
		return nil, err
	}

	return &producer, err
}

func (p *Producer) SendMessage(producer *sarama.AsyncProducer) {

	saramaMsg := &sarama.ProducerMessage{
		Topic: p.message.Topic,
		Value: sarama.ByteEncoder(p.message.Value),
	}

	var heds []sarama.RecordHeader
	for _, obj := range p.message.Headers {
		heds = append(heds, sarama.RecordHeader{
			Key:   []byte(obj.Key),
			Value: []byte(obj.Value),
		})
	}

	saramaMsg.Headers = heds
	(*producer).Input() <- saramaMsg

	go func() {
		for err := range (*producer).Errors() {
			if err != nil {
				p.callback(p.GetErrorMessage())
				fmt.Printf("Failed for message produced: %s \n", saramaMsg.Value)
			}

		}
	}()

}

func (p *Producer) GetErrorMessage() string {
	value := string(p.message.Value)

	return value
}
