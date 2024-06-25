

<h1>how to use aws sns and sqs?</h1>
<br />
<p>
install <a href="https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html" title="aws cli">aws cli</a> 
<br />

install docker container:<br /><br/>
sudo docker pull localstack/localstack<br />
sudo docker container run -it -d -p 4566:4566 localstack/localstack start<br />


<h1>How to configure the local environment?</h1><br />
<code>
aws configure<br />
	AWS Access Key ID [None]: fakeAccessKeyId <br />
	AWS Secret Access Key [None]: fakeSecretAccessKey<br />
	Default region name [us-east-1]: us-east-1<br />
	Default output format [None]: json<br />
    
</code>
aws configure --profile localstack<br />
	AWS Access Key ID [None]: nome_perfil_novo<br />
	AWS Secret Access Key [None]: senha_perfil_novo<br />
	Default region name [None]: us-east-1<br />
	Default output format [None]: json<br />
</p>
<p>
<h1>How to create topic and queue?</h1><br />

Create topic:<br/><br/>
aws --endpoint-url=http://localhost:4566 sns create-topic --name my-topic
<br /><br/>
Create queue
<br /><br/>
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue --region us-east-1
<br /><br/>
Create QueueArn
<br /><br/>
aws --endpoint-url=http://localhost:4566 sqs get-queue-attributes --queue-url http://localhost:4566/000000000000/my-queue --attribute-names QueueArn --region us-east-1
<br /><br/>
Subscribe SNS Topic to SQS Queue Endpoint
<br /><br/>
aws --endpoint-url=http://localhost:4566 sns subscribe --topic-arn $TOPIC_ARN --protocol sqs --notification-endpoint $QUEUE_ARN --region us-east-1
</p>