package main

import (
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/consumer"
)

func main() {
	go consumer.Consumer(&[]string{"springboot:9092"}, &[]string{"testar"}, "contacts")
	go consumer.ListenPartition(&[]string{"springboot:9092"}, "testar", 1, -1)
	select {}

}
