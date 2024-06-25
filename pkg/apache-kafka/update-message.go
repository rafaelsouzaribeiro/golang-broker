package apachekafka

import (
	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func UpdateKafkaMessage(msg *sarama.ConsumerMessage) *payload.Message {

	message := payload.Message{
		Value:     string(msg.Value),
		Topic:     msg.Topic,
		Partition: msg.Partition,
		Key:       string(msg.Key),
		Time:      msg.Timestamp,
	}

	var headers []payload.Header
	for _, header := range msg.Headers {
		headers = append(headers, payload.Header{
			Key:   string(header.Key),
			Value: string(header.Value),
		})
	}
	message.Headers = &headers

	return &message
}
