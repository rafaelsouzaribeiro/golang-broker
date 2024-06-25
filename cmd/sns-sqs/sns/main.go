package main

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/sns-sqs/sns"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/utils"
)

func main() {
	configs := utils.SNSSQSMessage{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
		Message:  "Message Test",
		TopicArn: "arn:aws:sns:us-east-1:000000000000:my-topic",
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		sns.Send(&configs)
		wg.Done()
	}()

	wg.Wait()

}
