package utils

import (
	"github.com/IBM/sarama"
)

func UpdateKafkaMessage(msg *sarama.ConsumerMessage) *Message {

	message := Message{
		Value:     string(msg.Value),
		Topic:     msg.Topic,
		Partition: msg.Partition,
		Key:       msg.Key,
		Time:      msg.Timestamp,
	}

	var headers []Header
	for _, header := range msg.Headers {
		headers = append(headers, Header{
			Key:   string(header.Key),
			Value: string(header.Value),
		})
	}
	message.Headers = headers

	return &message
}
