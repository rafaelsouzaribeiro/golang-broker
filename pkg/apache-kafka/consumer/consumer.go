package consumer

import (
	"log"

	"github.com/IBM/sarama"
	apachekafka "github.com/rafaelsouzaribeiro/golang-broker/pkg/apache-kafka"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

type ExampleConsumerGroupHandler struct {
	Channel chan<- payload.Message
}

func (*ExampleConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	log.Println("Consumer group is being rebalanced")
	return nil
}

func (*ExampleConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("Rebalancing will happen soon, current session will end")
	return nil
}

func (h *ExampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {

		data := apachekafka.UpdateKafkaMessage(msg)
		h.Channel <- *data
		sess.MarkMessage(msg, "")
	}

	return nil
}
