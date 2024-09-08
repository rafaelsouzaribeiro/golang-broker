package main

import (
	"sync"

	apachekafka "github.com/rafaelsouzaribeiro/golang-broker/pkg/apache-kafka"
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

	pro := factory.NewBroker(factory.Kafka, "springboot:9092")
	prod, err := pro.GetProducer(apachekafka.GetConfigProducer())

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := (*prod).Close(); err != nil {
			panic(err)
		}
	}()

	pro.SendMessage(prod, &message)

}
