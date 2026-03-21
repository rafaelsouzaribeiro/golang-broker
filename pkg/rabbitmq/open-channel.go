package rabbitmq

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMQBroker) OpenChannel() (*amqp.Channel, error) {
	con, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", os.Getenv("RABBITMQ_DEFAULT_USER"), os.Getenv("RABBITMQ_DEFAULT_PASS"), "localhost", 5672))
	if err != nil {
		return nil, err
	}
	ch, err := con.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}
