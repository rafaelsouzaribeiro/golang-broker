package main

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/consumer"
	"github.com/rafaelsouzaribeiro/apache-kafka/utils"
)

func main() {
	go Consumer(handleMessage)
	select {}

}

func Consumer(handleMessageFunc consumer.MessageCallback) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	msg := utils.Message{
		Topic:   "contact-adm-insert",
		GroupID: "contacts",
	}

	con := consumer.NewConsumer([]string{"springboot:9092"}, config, msg, handleMessageFunc)

	client, err := con.GetConsumer()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			panic(err)
		}

	}()

	cancel, err := con.VerifyConsumer(client)
	defer cancel()

	if err != nil {
		panic(err)
	}

	con.VerifyError(client)

}

func handleMessage(msgs utils.Message) {
	fmt.Printf("topic: %s, Message: %s, Partition: %d, Key: %d, time: %s\n", msgs.Topic, msgs.Value, msgs.Partition, msgs.Key, msgs.Time.Format("2006-01-02 15:04:05"))

	println("Headers:")
	for _, header := range msgs.Headers {
		fmt.Printf("Key: %s, Value: %s\n", header.Key, header.Value)
	}
}
