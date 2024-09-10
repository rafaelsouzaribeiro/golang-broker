package redis

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func (b *Broker) SendMessage(data *payload.Message) {

	messageByte, err := json.Marshal(data)
	if err != nil {
		log.Printf("error when running marshal %v", err)
	}

	b.client.LPush(context.Background(), data.Topic, messageByte)
}
