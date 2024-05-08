package consumer

import (
	"fmt"

	"github.com/IBM/sarama"
)

func ListenPartition(broker *[]string, topic string, partition int32, offset int64) {

	consumer, err := sarama.NewConsumer(*broker, GetConfig())

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
