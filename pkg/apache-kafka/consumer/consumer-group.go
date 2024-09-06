package consumer

import (
	"context"

	"github.com/IBM/sarama"
)

func (b *Broker) Consumer() {

	group, err := sarama.NewConsumerGroup(*b.broker, b.data.GroupID, GetConfig())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	handler := &ExampleConsumerGroupHandler{
		Channel: b.channel,
	}

	errs := group.Consume(ctx, *b.data.Topics, handler)
	if errs != nil {
		panic(errs)
	}

}
