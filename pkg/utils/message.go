package utils

import (
	"time"
)

type Message struct {
	Topics    *[]string
	Topic     string
	GroupID   string
	Value     string
	Partition int32
	Key       string
	Headers   *[]Header
	Time      time.Time
	Offset    int64
}

type Header struct {
	Key   string
	Value string
}

type SNSMessage struct {
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
