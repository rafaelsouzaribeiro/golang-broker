package main

import (
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/kafka/consumer"
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/utils"
)

func main() {
	data := utils.Message{
		Topics:    []string{"testar"},
		Topic:     "testar",
		GroupID:   "contacts",
		Partition: 1,
		Offset:    -1,
	}
	go consumer.Consumer(&[]string{"springboot:9092"}, &data)
	go consumer.ListenPartition(&[]string{"springboot:9092"}, &data)
	select {}

}
