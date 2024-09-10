package apachekafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func (p *Broker) SendMessage(data *payload.Message) {

	producer, err := GetProducer(GetConfigProducer(), p.broker)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := (*producer).Close(); err != nil {
			panic(err)
		}
	}()

	saramaMsg := &sarama.ProducerMessage{
		Topic: data.Topic,
		Value: sarama.ByteEncoder(data.Value),
	}

	if data.Headers != nil && len(*data.Headers) > 0 {
		var heds []sarama.RecordHeader
		for _, obj := range *data.Headers {
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
				fmt.Printf("Failed for message produced: %s \n", saramaMsg.Value)
			}

		}
	}()

}
