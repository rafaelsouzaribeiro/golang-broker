package consumer

import (
	"context"

	"github.com/IBM/sarama"
)

type MessageCallback func(messages []string)

type ConsumerGroupHandler struct {
	brokers  []string
	groupId  string
	topics   []string
	config   *sarama.Config
	callback MessageCallback
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

		messages := []string{string(msg.Value)}
		c.callback(messages)
		session.MarkMessage(msg, "") // Marca a mensagem como processada

	}

	return nil
}

func NewConsumer(brokers []string, groupId string, topics []string, config *sarama.Config, callback MessageCallback) *ConsumerGroupHandler {
	return &ConsumerGroupHandler{
		brokers:  brokers,
		groupId:  groupId,
		topics:   topics,
		config:   config,
		callback: callback,
	}
}

func (p *ConsumerGroupHandler) GetConsumer() (sarama.ConsumerGroup, error) {

	client, err := sarama.NewConsumerGroup(p.brokers, p.groupId, p.config)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (p *ConsumerGroupHandler) VerifyConsumer(client sarama.ConsumerGroup) (context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())

	handler := NewConsumer(p.brokers, p.groupId, p.topics, p.config, p.callback) // Atribua o manipulador retornado ao campo handler

	err := client.Consume(ctx, p.topics, handler)
	if err != nil {
		return cancel, err
	}

	return cancel, nil

}

func (p *ConsumerGroupHandler) VerifyError(client sarama.ConsumerGroup) error {

	for err := range client.Errors() {
		if err != nil {
			return err
		}

	}

	// go func() {
	// 	for {
	// 		err := <-client.Errors()
	// 		if err != nil {
	// 			// Se houver um erro, feche o canal de erros e encerre a goroutine
	// 			fmt.Println("Ocorreu algum erro")

	// 			p.errors <- err
	// 			close(p.errors)
	// 			return
	// 		}
	// 	}
	// }()

	// for err := range p.errors {
	// 	fmt.Println(err.Error())
	// 	return err // Assuming you want to stop after receiving the first error
	// }

	return nil
}
