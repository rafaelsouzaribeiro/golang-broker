package producer

import (
	"fmt"

	"github.com/IBM/sarama"
)

func (p *Producer) SendMessage(producer *sarama.AsyncProducer) {

	saramaMsg := &sarama.ProducerMessage{
		Topic: p.message.Topic,
		Value: sarama.ByteEncoder(p.message.Value),
	}

	if p.message.Headers != nil && len(*p.message.Headers) > 0 {
		var heds []sarama.RecordHeader
		for _, obj := range *p.message.Headers {
			heds = append(heds, sarama.RecordHeader{
				Key:   []byte(obj.Key),
				Value: []byte(obj.Value),
			})
		}

		saramaMsg.Headers = heds
	}

	(*producer).Input() <- saramaMsg

	go func() {
		for err := range (*producer).Errors() {
			if err != nil {
				p.callback(*p.GetErrorMessage())
				fmt.Printf("Failed for message produced: %s \n", saramaMsg.Value)
			}

		}
	}()

}
