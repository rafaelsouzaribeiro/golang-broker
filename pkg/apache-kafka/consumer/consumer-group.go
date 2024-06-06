package consumer

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/utils"
)

func Consumer(broker *[]string, data *utils.Message, callback MessageCallback) {

	group, err := sarama.NewConsumerGroup(*broker, data.GroupID, GetConfig())
	if err != nil {
		panic(err)
	}

	// Consume messages
	ctx := context.Background()

	handler := &ExampleConsumerGroupHandler{
		Callback: callback,
	}

	errs := group.Consume(ctx, *data.Topics, handler)
	if errs != nil {
		panic(errs)
	}

}
