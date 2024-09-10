package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func (b *Broker) Consumer(data *payload.Message, channel chan<- payload.Message) {

	for {
		res, err := b.client.BLPop(context.Background(), 0*time.Second, *data.Topics...).Result()
		if err != nil {
			fmt.Println("Error reading item from queue:", err)
			continue
		}

		var tmp payload.Message

		json.Unmarshal([]byte(res[1]), &tmp)

		channel <- tmp

	}
}
