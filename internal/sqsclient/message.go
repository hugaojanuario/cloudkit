package sqsclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func SendMessage(ctx context.Context, client *sqs.Client, queueURL string, message string) (string, error) {
	input := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(message),
	}

	output, err := client.SendMessage(ctx, input)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar a mensagem: %w", err)
	}
	return aws.ToString(output.MessageId), nil
}
