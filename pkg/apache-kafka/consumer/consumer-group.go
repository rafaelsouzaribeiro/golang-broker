package consumer

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/utils"
)

func Consumer(broker *[]string, data *utils.Message, channel chan utils.Message) {

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
