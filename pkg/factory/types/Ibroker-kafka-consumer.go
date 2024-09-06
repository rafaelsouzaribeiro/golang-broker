package types

type IbrokerKafkaConsumer interface {
	Consumer()
	ListenPartition()
}
