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
	// Configurar sessão AWS com o LocalStack
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
	}))

	// Criar um cliente SQS
	svc := sqs.New(sess)

	// Parâmetros para receber mensagens da fila SQS
	receiveMessageInput := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(10), // Número máximo de mensagens a serem recebidas
		WaitTimeSeconds:     aws.Int64(20), // Tempo de espera (long polling)
	}

	// Receber mensagens da fila SQS
	result, err := svc.ReceiveMessage(receiveMessageInput)
	if err != nil {
		log.Fatalf("Erro ao receber mensagens: %v", err)
	}

	// Exibir as mensagens recebidas
	if len(result.Messages) > 0 {
		for _, message := range result.Messages {
			fmt.Printf("Mensagem recebida: %s\n", *message.Body)

			// Apagar a mensagem após o processamento
			deleteMessageInput := &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: message.ReceiptHandle,
			}
			_, err := svc.DeleteMessage(deleteMessageInput)
			if err != nil {
				log.Fatalf("Erro ao apagar a mensagem: %v", err)
			}
		}
	} else {
		fmt.Println("Nenhuma mensagem recebida.")
	}
}
