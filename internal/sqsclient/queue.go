package sqsclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func CreateQueue(ctx context.Context, client *sqs.Client, queueName string) (string, error) {
	output, err := client.CreateQueue(ctx, &sqs.CreateQueueInput{QueueName: aws.String(queueName)})
	if err != nil {
		return "", fmt.Errorf("criar fila: %w", err)
	}
	return aws.ToString(output.QueueUrl), nil
}
