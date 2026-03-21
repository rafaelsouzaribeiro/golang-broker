package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func (r *RabbitMQBroker) Publish(payload *payload.RabbitMQMessage) error {
	err := payload.Channel.Publish(
		payload.Exchange,
		payload.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        payload.Value,
			Headers: amqp.Table{
				"header": payload.Header,
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
