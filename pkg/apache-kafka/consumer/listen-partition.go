package consumer

import (
	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func (b *Broker) ListenPartition() {

	consumer, err := sarama.NewConsumer(*b.broker, GetConfig())

	if err != nil {
		panic(err)
	}

	pc, err := consumer.ConsumePartition(b.data.Topic, b.data.Partition, b.data.Offset)

	if err != nil {
		panic(err)
	}

	for msgs := range pc.Messages() {
		var listHeaders []payload.Header

		for _, h := range msgs.Headers {
			header := payload.Header{Key: string(h.Key), Value: string(h.Value)}
			listHeaders = append(listHeaders, header)
		}

		b.channel <- payload.Message{
			Topic:     msgs.Topic,
			GroupID:   b.data.GroupID,
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
