package main

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/consumer"
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/utils"
)

func main() {
	//go Consumer(&[]string{"springboot:9092"}, &[]string{"testar"}, "contacts")
	go ListenPartition(&[]string{"springboot:9092"}, "testar", 1, -1)
	select {}

}

func ListenPartition(broker *[]string, topic string, partition int32, offset int64) {

	consumer, err := sarama.NewConsumer(*broker, consumer.GetConfig())

	if err != nil {
		panic(err)
	}

	pc, err := consumer.ConsumePartition(topic, partition, offset)

	if err != nil {
		panic(err)
	}
	for msgs := range pc.Messages() {
		fmt.Printf("topic: %s, Message: %s, Partition: %d, Key: %s, time: %s\n", msgs.Topic, msgs.Value, msgs.Partition, msgs.Key, msgs.Timestamp.Format("2006-01-02 15:04:05"))

		println("Headers:")
		for _, header := range msgs.Headers {
			fmt.Printf("Key: %s, Value: %s\n", header.Key, header.Value)
		}

	}

	pc.Close()

}

func Consumer(broker, topics *[]string, groupId string) {

	group, err := sarama.NewConsumerGroup(*broker, groupId, consumer.GetConfig())
	if err != nil {
		panic(err)
	}

	// Consume messages
	ctx := context.Background()

	handler := &consumer.ExampleConsumerGroupHandler{
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
