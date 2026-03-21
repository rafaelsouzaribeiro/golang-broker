package rabbitmq

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func (r *RabbitMQBroker) Consume(payload *payload.RabbitMQMessage) error {

	msgs, err := payload.Channel.Consume(
		payload.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		payload.Msgs <- &msg
	}

	return nil
}
