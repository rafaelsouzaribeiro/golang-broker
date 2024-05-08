package kafka

import (
	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/broker-golang/pkg/utils"
)

func UpdateKafkaMessage(msg *sarama.ConsumerMessage) *utils.Message {

	message := utils.Message{
		Value:     string(msg.Value),
		Topic:     msg.Topic,
		Partition: msg.Partition,
		Key:       string(msg.Key),
		Time:      msg.Timestamp,
	}

	var headers []utils.Header
	for _, header := range msg.Headers {
		headers = append(headers, utils.Header{
			Key:   string(header.Key),
			Value: string(header.Value),
		})
	}
	message.Headers = headers

	return &message
}
