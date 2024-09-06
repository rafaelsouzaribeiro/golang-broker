package main

import (
	"fmt"
	"sync"

	"github.com/rafaelsouzaribeiro/golang-broker/pkg/apache-kafka/producer"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		Producer()
		wg.Done()
	}()

	wg.Wait()
}

func Producer() {

	message := payload.Message{
		Value: "Testar",
		Topic: "contact-adm-insert",
		Headers: &[]payload.Header{
			{
				Key:   "your-header-key1",
				Value: "your-header-value1",
			},
			{
				Key:   "your-header-key2",
				Value: "your-header-value2",
			},
		},
	}

	pro := factory.IProducerBroker(factory.APACHE_KAFKA_PRODUCER, &[]string{"springboot:9092"}, &message, producer.GetConfig(), func(messages payload.Message) {
		fmt.Printf("message failure: %s, topic failure: %s, partition failure: %d \n", messages.Value, messages.Topic, messages.Partition)
	})

	prod, err := pro.GetProducer()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := (*prod).Close(); err != nil {
			panic(err)
		}
	}()

	pro.SendMessage(prod)

}
