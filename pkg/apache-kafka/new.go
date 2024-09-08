package apachekafka

// type Broker struct {
// 	broker  *[]string
// 	data    *payload.Message
// 	channel chan<- payload.Message
// }

type Broker struct {
	broker string
}

func NewBroker(broker string) *Broker {
	return &Broker{
		broker: broker,
	}
}
