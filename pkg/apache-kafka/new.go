package apachekafka

type Broker struct {
	broker string
}

func NewBroker(broker string) *Broker {
	return &Broker{
		broker: broker,
	}
}
