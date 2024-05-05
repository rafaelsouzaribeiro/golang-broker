package consumer

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/utils"
)

type MessageCallback func(messages utils.Message)

type ConsumerGroupHandler struct {
	brokers  []string
	config   *sarama.Config
	callback MessageCallback
	message  utils.Message
}

// Cleanup implements sarama.ConsumerGroupHandler.
func (c *ConsumerGroupHandler) Cleanup(s sarama.ConsumerGroupSession) error {
	return nil
}

// Setup implements sarama.ConsumerGroupHandler.
func (c *ConsumerGroupHandler) Setup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {

		datas := utils.UpdateKafkaMessage(msg)
		c.callback(*datas)
		session.MarkMessage(msg, "") // Marca a mensagem como processada

	}

	return nil
}

func NewConsumer(brokers []string, config *sarama.Config, data utils.Message, callback MessageCallback) *ConsumerGroupHandler {
	return &ConsumerGroupHandler{
		brokers:  brokers,
		config:   config,
		callback: callback,
		message:  data,
	}
}

func (p *ConsumerGroupHandler) GetConsumer() (sarama.ConsumerGroup, error) {

	client, err := sarama.NewConsumerGroup(p.brokers, p.message.GroupID, p.config)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (p *ConsumerGroupHandler) VerifyConsumer(client sarama.ConsumerGroup) (context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())

	handler := NewConsumer(p.brokers, p.config, utils.Message{}, p.callback) // Atribua o manipulador retornado ao campo handler

	err := client.Consume(ctx, []string{p.message.Topic}, handler)
	if err != nil {
		return cancel, err
	}

	return cancel, nil

}

func (p *ConsumerGroupHandler) VerifyError(client sarama.ConsumerGroup) {

	go func() {
		for err := range client.Errors() {
			if err != nil {
				fmt.Printf("consumer error: %s \n", err)
			}

		}
	}()
}
