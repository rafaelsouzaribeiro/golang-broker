package redis

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory/types"
	"github.com/redis/go-redis/v9"
)

type Broker struct {
	client *redis.Client
}

func NewBroker(broker string, password string) types.Ibroker {
	rdb := redis.NewClient(&redis.Options{
		Addr:     broker,
		Password: password,
	})

	return &Broker{
		client: rdb,
	}

}
