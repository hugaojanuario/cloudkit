package awsconfig

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func Load(ctx context.Context) (aws.Config, error) {

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("mock-id", "mock-key", "")),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return aws.Config{}, fmt.Errorf("carregando aws config: %w", err)
	}

	return cfg, nil
}
