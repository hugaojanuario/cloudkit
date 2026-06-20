package sqsclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func Delete(ctx context.Context, client *sqs.Client, queueURL string, receiptHandle string) error {
	input := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	}

	_, err := client.DeleteMessage(ctx, input)
	if err != nil {
		return fmt.Errorf("erro ao deletar mensagem: %w", err)
	}
	return nil
}
