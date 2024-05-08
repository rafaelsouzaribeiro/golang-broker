package utils

import (
	"time"
)

type Message struct {
	Topics    []string
	Topic     string
	GroupID   string
	Value     string
	Partition int32
	Key       string
	Headers   []Header
	Time      time.Time
	Offset    int64
}

type Header struct {
	Key   string
	Value string
}
