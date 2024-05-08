package consumer

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/utils"
)

func Consumer(broker, topics *[]string, groupId string) {

	group, err := sarama.NewConsumerGroup(*broker, groupId, GetConfig())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	handler := &ExampleConsumerGroupHandler{
		Callback: handleMessage,
	}

	errs := group.Consume(ctx, *topics, handler)
	if errs != nil {
		panic(errs)
	}

}

func handleMessage(msgs utils.Message) {
	fmt.Printf("topic: %s, Message: %s, Partition: %d, Key: %d, time: %s\n", msgs.Topic, msgs.Value, msgs.Partition, msgs.Key, msgs.Time.Format("2006-01-02 15:04:05"))

	println("Headers:")
	for _, header := range msgs.Headers {
		fmt.Printf("Key: %s, Value: %s\n", header.Key, header.Value)
	}
}
