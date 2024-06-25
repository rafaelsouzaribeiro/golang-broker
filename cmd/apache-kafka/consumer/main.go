package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/golang-broker/pkg/apache-kafka/consumer"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func main() {

	data := payload.Message{
		Topics:    &[]string{"contact-adm-insert", "testar"},
		Topic:     "contact-adm-insert",
		GroupID:   "contacts",
		Partition: 0,
		Offset:    -1,
	}
	canal := make(chan payload.Message)
	go consumer.Consumer(&[]string{"springboot:9092"}, &data, canal)
	go consumer.ListenPartition(&[]string{"springboot:9092"}, &data, canal)

	for msgs := range canal {
		printMessage(&msgs)
	}

	close(canal)

	select {}

}

func printMessage(msgs *payload.Message) {
	fmt.Printf("topic: %s, Message: %s, Partition: %d, Key: %s, time: %s\n", msgs.Topic, msgs.Value, msgs.Partition, msgs.Key, msgs.Time.Format("2006-01-02 15:04:05"))

	println("Headers:")
	for _, header := range *msgs.Headers {
		fmt.Printf("Key: %s, Value: %s\n", header.Key, header.Value)
	}
}
