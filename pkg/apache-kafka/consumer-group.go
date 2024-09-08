package apachekafka

import (
	"context"
	"strings"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func (b *Broker) Consumer(data *payload.Message, channel chan<- payload.Message) {

	broker := strings.Split(b.broker, ",")
	group, err := sarama.NewConsumerGroup(broker, data.GroupID, GetConfigConsumer())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	handler := &ExampleConsumerGroupHandler{
		Channel: channel,
	}

	errs := group.Consume(ctx, *data.Topics, handler)
	if errs != nil {
		panic(errs)
	}

}
