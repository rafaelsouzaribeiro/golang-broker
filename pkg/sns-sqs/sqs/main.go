package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	queueURL = "http://localhost:4566/000000000000/my-queue"
)

func main() {
	go Sqs()
	select {}
}

func Sqs() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
	}))

	svc := sqs.New(sess)

	for {

		receiveMessageInput := &sqs.ReceiveMessageInput{
			MaxNumberOfMessages: aws.Int64(1),
			QueueUrl:            aws.String(queueURL),
			WaitTimeSeconds:     aws.Int64(20),
		}

		result, err := svc.ReceiveMessage(receiveMessageInput)
		if err != nil {
			log.Fatalf("Error receiving messages: %v", err)
			continue
		}

		if len(result.Messages) > 0 {
			for _, message := range result.Messages {
				fmt.Printf("Message received: %s\n", *message.Body)

				deleteMessageInput := &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(queueURL),
					ReceiptHandle: message.ReceiptHandle,
				}
				_, err := svc.DeleteMessage(deleteMessageInput)
				if err != nil {
					log.Fatalf("Error deleting message: %v", err)
					continue
				}
			}
		}

	}

}
