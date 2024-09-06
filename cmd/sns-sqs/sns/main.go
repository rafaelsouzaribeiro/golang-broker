package main

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func main() {
	configs := payload.SNSSQSMessage{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
		Message:  "Message Test",
		TopicArn: "arn:aws:sns:us-east-1:000000000000:my-topic",
	}

	var wg sync.WaitGroup
	wg.Add(1)

	factory := factory.ISNSBroker(factory.SNS, &configs)

	go func() {
		factory.Send()
		wg.Done()
	}()

	wg.Wait()

}
