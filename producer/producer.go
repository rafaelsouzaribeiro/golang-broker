package producer

import (
	"fmt"

	"github.com/IBM/sarama"
)

type MessageCallback func(messages string)

type Producer struct {
	addrs    []string
	topic    string
	message  sarama.Encoder
	config   *sarama.Config
	callback MessageCallback
}

func NewProducer(addrs []string, topic string, message sarama.Encoder, config *sarama.Config, callback MessageCallback) *Producer {
	return &Producer{
		addrs:    addrs,
		topic:    topic,
		message:  message,
		config:   config,
		callback: callback,
	}
}

func (p *Producer) GetProducer() (*sarama.AsyncProducer, error) {

	producer, err := sarama.NewAsyncProducer(p.addrs, p.config)

	if err != nil {
		p.callback(p.GetErrorMessage())
		return nil, err
	}

	return &producer, err
}

func (p *Producer) SendMessage(producer *sarama.AsyncProducer) {

	message := &sarama.ProducerMessage{Topic: p.topic, Value: p.message}

	(*producer).Input() <- message

	go func() {
		for err := range (*producer).Errors() {
			if err != nil {
				p.callback(p.GetErrorMessage())
				fmt.Printf("Failed for message produced: %s \n", message.Value)
			}

		}
	}()

	// select {
	// case success := <-(*producer).Successes():
	// 	value, err := success.Value.Encode()
	// 	if err != nil {
	// 		fmt.Println("Error encoding message:", err)
	// 		return err
	// 	}

	// 	fmt.Println("Mensagem produzida:", string(value))
	// case err := <-(*producer).Errors():
	// 	fmt.Println("Falho para mensagem produzida:", err)
	// 	return err
	// }

	// return nil
}

func (p *Producer) GetErrorMessage() string {
	value, _ := p.message.Encode()

	return string(value)
}
