package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	topicArn = "arn:aws:sns:us-east-1:000000000000:my-topic"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
	}))

	svc := sns.New(sess)
	message := "Message Test"

	publishParams := &sns.PublishInput{
		Message:  aws.String("Message Test"),
		TopicArn: aws.String(topicArn),
	}

	result, err := svc.Publish(publishParams)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Message sent successfully:\nMessage ID: %s\nMessage Text: %s\n", *result.MessageId, message)
}
