package consumer

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/utils"
)

type MessageCallback func(messages utils.Message)

type ExampleConsumerGroupHandler struct {
	Callback MessageCallback
}

func (*ExampleConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	// Mark the beginning of a new session. This is called when the consumer group is being rebalanced.
	log.Println("Consumer group is being rebalanced")
	return nil
}

func (*ExampleConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	// Mark the end of the current session. This is called just before the next rebalance happens.
	log.Println("Rebalancing will happen soon, current session will end")
	return nil
}

func (h *ExampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// This is where you put your message handling logic
	for msg := range claim.Messages() {

		data := utils.UpdateKafkaMessage(msg)
		h.Callback(*data)

		sess.MarkMessage(msg, "")
	}

	return nil
}
