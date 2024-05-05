package producer

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/pkg/utils"
)

type MessageCallback func(messages string)

type Producer struct {
	addrs    []string
	topic    string
	message  []byte
	config   *sarama.Config
	callback MessageCallback
}

func NewProducer(addrs []string, topic string, message []byte, config *sarama.Config, callback MessageCallback) *Producer {
	return &Producer{
		addrs:    addrs,
		topic:    topic,
		message:  message,
		config:   config,
		callback: callback,
	}
}

func Encode(data utils.Message) ([]byte, error) {
	// Encode the message as JSON
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte) (*utils.Message, error) {
	// Create a new message instance
	message := &utils.Message{}

	// Decode the JSON data
	decoder := json.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(message)
	if err != nil {
		return nil, err
	}

	return message, nil
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

	msg, err := Decode(p.message)

	if err != nil {
		panic(err)
	}

	saramaMsg := &sarama.ProducerMessage{
		Topic: msg.Topic,
		Value: sarama.ByteEncoder(msg.Value),
	}

	var heds []sarama.RecordHeader
	for _, obj := range msg.Headers {
		heds = append(heds, sarama.RecordHeader{
			Key:   []byte(obj.Key),
			Value: []byte(obj.Value),
		})
	}

	saramaMsg.Headers = heds
	(*producer).Input() <- saramaMsg

	go func() {
		for err := range (*producer).Errors() {
			if err != nil {
				p.callback(p.GetErrorMessage())
				fmt.Printf("Failed for message produced: %s \n", saramaMsg.Value)
			}

		}
	}()

}

func (p *Producer) GetErrorMessage() string {
	value := string(p.message)

	return string(value)
}
