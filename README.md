

<h1>how to use aws sns and sqs?</h1>
<br />
<p>
install <a href="https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html" title="aws cli">aws cli</a> 
<br />

install docker container:<br />
sudo docker pull localstack/localstack<br />
sudo docker container run -it -d -p 4566:4566 localstack/localstack start<br />

aws configure<br />
	AWS Access Key ID [None]: fakeAccessKeyId 
	AWS Secret Access Key [None]: fakeSecretAccessKey
	Default region name [us-east-1]: us-east-1
	Default output format [None]: json
    

aws configure --profile localstack<br />
	AWS Access Key ID [None]: nome_perfil_novo
	AWS Secret Access Key [None]: senha_perfil_novo
	Default region name [None]: us-east-1
	Default output format [None]: json
</p>
<p>
Create topic:<br/>
aws --endpoint-url=http://localhost:4566 sns create-topic --name my-topic
<br />
Create queue
<br />
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue --region us-east-1
<br />
Create QueueArn
<br />
aws --endpoint-url=http://localhost:4566 sqs get-queue-attributes --queue-url http://localhost:4566/000000000000/my-queue --attribute-names QueueArn --region us-east-1
<br />
Subscribe SNS Topic to SQS Queue Endpoint
<br />
aws --endpoint-url=http://localhost:4566 sns subscribe --topic-arn $TOPIC_ARN --protocol sqs --notification-endpoint $QUEUE_ARN --region us-east-1
</p>