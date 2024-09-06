package sns

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func (b *Broker) Send() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: b.Config.Endpoint,
		Region:   b.Config.Region,
	}))

	svc := sns.New(sess)
	message := b.Config.Message

	publishParams := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(b.Config.TopicArn),
	}

	result, err := svc.Publish(publishParams)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Message sent successfully:\nMessage ID: %s\nMessage Text: %s\n", *result.MessageId, message)
}
