package sqsclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func ReceiveMessage(ctx context.Context, client *sqs.Client, queueUrl string) ([]types.Message, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     20,
		VisibilityTimeout:   30,
	}

	output, err := client.ReceiveMessage(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("erro ao receber mensagem: %w", err)
	}

	return output.Messages, nil
}
