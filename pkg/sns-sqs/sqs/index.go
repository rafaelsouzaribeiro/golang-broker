package sqs

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func (b *Broker) Receive() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: b.Configs.Endpoint,
		Region:   b.Configs.Region,
	}))

	svc := sqs.New(sess)

	for {

		receiveMessageInput := &sqs.ReceiveMessageInput{
			MaxNumberOfMessages: aws.Int64(1),
			QueueUrl:            aws.String(b.Configs.QueueURL),
			WaitTimeSeconds:     aws.Int64(10),
		}

		result, err := svc.ReceiveMessage(receiveMessageInput)
		if err != nil {
			log.Fatalf("Error receiving messages: %v", err)
			continue
		}

		if len(result.Messages) > 0 {
			for _, message := range result.Messages {
				var snsMessage payload.SNSSQSMessage

				err := json.Unmarshal([]byte(*message.Body), &snsMessage)
				if err != nil {
					log.Printf("Error decoding message: %v", err)
				}

				b.MessageChan <- snsMessage

				deleteMessageInput := &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(b.Configs.QueueURL),
					ReceiptHandle: message.ReceiptHandle,
				}
				_, errs := svc.DeleteMessage(deleteMessageInput)
				if errs != nil {
					log.Fatalf("Error deleting message: %v", errs)
					continue
				}
			}
		}

	}

}
