package producer

// type MessageCallback func(messages payload.Message)

// type Producer struct {
// 	addrs    []string
// 	message  payload.Message
// 	config   *sarama.Config
// 	callback MessageCallback
// }

// func NewProducer(addrs *[]string, message *payload.Message, config *sarama.Config, callback MessageCallback) types.IbrokerKafkaProducer {
// 	return &Producer{
// 		addrs:    *addrs,
// 		message:  *message,
// 		config:   config,
// 		callback: callback,
// 	}
// }

type Broker struct {
	broker string
}

func NewBroker(broker string) *Broker {
	return &Broker{
		broker: broker,
	}
}
