package sns

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/rafaelsouzaribeiro/broker-golang/pkg/utils"
)

func Sns(config utils.SNSMessage) {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: config.Endpoint,
		Region:   config.Region,
	}))

	svc := sns.New(sess)
	message := config.Message

	publishParams := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(config.TopicArn),
	}

	result, err := svc.Publish(publishParams)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Message sent successfully:\nMessage ID: %s\nMessage Text: %s\n", *result.MessageId, message)
}
