package sqs

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/utils"
)

func Sqs(configs utils.SNSSQSMessage, messageChan chan<- utils.SNSSQSMessage) {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: configs.Endpoint,
		Region:   configs.Region,
	}))

	svc := sqs.New(sess)

	for {

		receiveMessageInput := &sqs.ReceiveMessageInput{
			MaxNumberOfMessages: aws.Int64(1),
			QueueUrl:            aws.String(configs.QueueURL),
			WaitTimeSeconds:     aws.Int64(10),
		}

		result, err := svc.ReceiveMessage(receiveMessageInput)
		if err != nil {
			log.Fatalf("Error receiving messages: %v", err)
			continue
		}

		if len(result.Messages) > 0 {
			for _, message := range result.Messages {
				var snsMessage utils.SNSSQSMessage

				err := json.Unmarshal([]byte(*message.Body), &snsMessage)
				if err != nil {
					log.Printf("Error decoding message: %v", err)
				}

				messageChan <- snsMessage

				deleteMessageInput := &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(configs.QueueURL),
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
