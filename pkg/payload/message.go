package payload

import (
	"time"
)

type Message struct {
	Topics    *[]string `json:"topics"`
	Topic     string    `json:"topic"`
	GroupID   string    `json:"group_id"`
	Value     []byte    `json:"value"`
	Partition int32     `json:"partition"`
	Key       []byte    `json:"key"`
	Headers   *[]Header `json:"header"`
	Time      time.Time `json:"time"`
	Offset    int64     `json:"offset"`
}

type Header struct {
	Key   string
	Value string
}

type SNSSQSMessage struct {
	Type             string `json:"Type"`
	MessageId        string `json:"MessageId"`
	TopicArn         string `json:"TopicArn"`
	Subject          string `json:"Subject"`
	Message          string `json:"Message"`
	Timestamp        string `json:"Timestamp"`
	SignatureVersion string `json:"SignatureVersion"`
	Endpoint         *string
	Region           *string
	QueueURL         string
}
