package apachekafka

import (
	"strings"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func (b *Broker) ListenPartition(data *payload.Message, channel chan<- payload.Message) {

	broker := strings.Split(b.broker, ",")
	consumer, err := sarama.NewConsumer(broker, GetConfigConsumer())

	if err != nil {
		panic(err)
	}

	pc, err := consumer.ConsumePartition(data.Topic, data.Partition, data.Offset)

	if err != nil {
		panic(err)
	}

	for msgs := range pc.Messages() {
		var listHeaders []payload.Header

		for _, h := range msgs.Headers {
			header := payload.Header{Key: string(h.Key), Value: string(h.Value)}
			listHeaders = append(listHeaders, header)
		}

		channel <- payload.Message{
			Topic:     msgs.Topic,
			GroupID:   data.GroupID,
			Value:     string(msgs.Value),
			Key:       string(msgs.Key),
			Partition: msgs.Partition,
			Headers:   &listHeaders,
			Time:      msgs.Timestamp,
			Offset:    msgs.Offset,
		}

	}

	pc.Close()

}
