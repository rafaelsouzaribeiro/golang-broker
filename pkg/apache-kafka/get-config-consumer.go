package apachekafka

import "github.com/IBM/sarama"

func GetConfigConsumer() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.Retry.Max = 5

	return config
}
