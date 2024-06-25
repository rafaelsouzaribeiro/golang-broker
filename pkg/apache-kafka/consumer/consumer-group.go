package consumer

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func Consumer(broker *[]string, data *payload.Message, channel chan payload.Message) {

	group, err := sarama.NewConsumerGroup(*broker, data.GroupID, GetConfig())
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
