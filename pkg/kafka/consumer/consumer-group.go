package consumer

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/broker-golang/pkg/utils"
)

func Consumer(broker *[]string, data *utils.Message) {

	group, err := sarama.NewConsumerGroup(*broker, data.GroupID, GetConfig())
	if err != nil {
		panic(err)
	}

	// Consume messages
	ctx := context.Background()

	handler := &ExampleConsumerGroupHandler{
		Callback: handleMessage,
	}

	errs := group.Consume(ctx, data.Topics, handler)
	if errs != nil {
		panic(errs)
	}

}

func handleMessage(msgs utils.Message) {
	fmt.Printf("topic: %s, Message: %s, Partition: %d, Key: %s, time: %s\n", msgs.Topic, msgs.Value, msgs.Partition, msgs.Key, msgs.Time.Format("2006-01-02 15:04:05"))

	println("Headers:")
	for _, header := range msgs.Headers {
		fmt.Printf("Key: %s, Value: %s\n", header.Key, header.Value)
	}
}
