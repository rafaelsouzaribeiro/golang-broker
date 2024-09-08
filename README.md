
<h1>how to use aws sns and sqs?</h1>
<br />
<p>
install <a href="https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html" title="aws cli">aws cli</a> 
<br />

run docker container:<br /><br/>
`sudo docker pull localstack/localstack`<br />
`sudo docker container run -it -d -p 4566:4566 localstack/localstack start`<br />


<h1>How can I set up the local development environment?</h1><br />

`aws configure`<br />
	AWS Access Key ID [None]: fakeAccessKeyId <br />
	AWS Secret Access Key [None]: fakeSecretAccessKey<br />
	Default region name [us-east-1]: us-east-1<br />
	Default output format [None]: json<br />
    

`aws configure --profile localstack`<br />
	AWS Access Key ID [None]: nome_perfil_novo<br />
	AWS Secret Access Key [None]: senha_perfil_novo<br />
	Default region name [None]: us-east-1<br />
	Default output format [None]: json<br />
</p>
<p>
<h1>How to create topic and queue?</h1><br />

Create topic:<br/><br/>
`aws --endpoint-url=http://localhost:4566 sns create-topic --name my-topic`
<br /><br/>
Create queue
<br /><br/>
`aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue --region us-east-1`
<br /><br/>
Create QueueArn
<br /><br/>
`aws --endpoint-url=http://localhost:4566 sqs get-queue-attributes --queue-url http://localhost:4566/000000000000/my-queue --attribute-names QueueArn --region us-east-1`
<br /><br/>
Subscribe SNS Topic to SQS Queue Endpoint
<br /><br/>
`aws --endpoint-url=http://localhost:4566 sns subscribe --topic-arn $TOPIC_ARN --protocol sqs --notification-endpoint $QUEUE_ARN --region us-east-1`
</p>

<h1>To use SQN and SNS, follow the code below:</h1>
Sqs:

```go
configs := payload.SNSSQSMessage{
	Endpoint: aws.String("http://localhost:4566"),
	Region:   aws.String("us-east-1"),
	QueueURL: "http://localhost:4566/000000000000/my-queue",
}

messageChan := make(chan payload.SNSSQSMessage)

factory := factory.ISQSSNSBroker(&configs)
go factory.Receive(messageChan)

for message := range messageChan {
	fmt.Printf("Received message: %s Message Id: %s Topic: %s Time: %s\n",
		message.Message, message.MessageId, message.TopicArn, message.Timestamp)
}

select {}
	
```
SNS:
```go
configs := payload.SNSSQSMessage{
	Endpoint: aws.String("http://localhost:4566"),
	Region:   aws.String("us-east-1"),
	Message:  "Message Test",
	TopicArn: "arn:aws:sns:us-east-1:000000000000:my-topic",
}

var wg sync.WaitGroup
wg.Add(1)

factory := factory.ISQSSNSBroker(&configs)

go func() {
	factory.Send()
	wg.Done()
}()

wg.Wait()
```
<h1>To use apache Kafka, follow the code below:</h1>
<br/>
Consumer:

```go
func main() {

	data := payload.Message{
		Topics:    &[]string{"contact-adm-insert", "testar"},
		Topic:     "contact-adm-insert",
		GroupID:   "contacts",
		Partition: 0,
		Offset:    -1,
	}
	canal := make(chan payload.Message)
	broker := factory.NewBroker(factory.Kafka, "springboot:9092")
	go broker.Consumer(&data, canal)
	go broker.ListenPartition(&data, canal)

	for msgs := range canal {
		printMessage(&msgs)
	}

	close(canal)

	select {}

}

func printMessage(msgs *payload.Message) {
	fmt.Printf("topic: %s, Message: %s, Partition: %d, Key: %s, time: %s\n", msgs.Topic, msgs.Value, msgs.Partition, msgs.Key, msgs.Time.Format("2006-01-02 15:04:05"))

	println("Headers:")
	for _, header := range *msgs.Headers {
		fmt.Printf("Key: %s, Value: %s\n", header.Key, header.Value)
	}
}

```

Producer:

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		Producer()
		wg.Done()
	}()

	wg.Wait()
}

func Producer() {

	message := payload.Message{
		Value: "Testar",
		Topic: "contact-adm-insert",
		Headers: &[]payload.Header{
			{
				Key:   "your-header-key1",
				Value: "your-header-value1",
			},
			{
				Key:   "your-header-key2",
				Value: "your-header-value2",
			},
		},
	}

	pro := factory.NewBroker(factory.Kafka, "springboot:9092")
	prod, err := pro.GetProducer(apachekafka.GetConfigProducer())

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := (*prod).Close(); err != nil {
			panic(err)
		}
	}()

	pro.SendMessage(prod, &message)

}

```