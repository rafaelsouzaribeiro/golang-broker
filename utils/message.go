package utils

import "time"

type Message struct {
	Topic     string
	GroupID   string
	Value     string
	Partition int32
	Key       []byte
	Headers   []Header
	Time      time.Time
}

type Header struct {
	Key   string
	Value string
}
