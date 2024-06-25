package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/sns-sqs/sqs"
)

func main() {
	configs := payload.SNSSQSMessage{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
		QueueURL: "http://localhost:4566/000000000000/my-queue",
	}

	messageChan := make(chan payload.SNSSQSMessage)

	go sqs.Receive(&configs, messageChan)

	for message := range messageChan {
		fmt.Printf("Received message: %s Message Id: %s Topic: %s Time: %s\n",
			message.Message, message.MessageId, message.TopicArn, message.Timestamp)
	}

	select {}
}
