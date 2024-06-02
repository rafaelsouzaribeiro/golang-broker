package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/rafaelsouzaribeiro/broker-golang/pkg/utils"
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
				var snsMessage utils.SNSMessage

				err := json.Unmarshal([]byte(*message.Body), &snsMessage)
				if err != nil {
					log.Printf("Erro ao decodificar a mensagem: %v", err)
				}

				fmt.Printf("Message received: Value:%s MessageId:%s Topic:%s Timestamp: %s\n",
					snsMessage.Message, snsMessage.MessageId, snsMessage.TopicArn, snsMessage.Timestamp)

				deleteMessageInput := &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(queueURL),
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
