package s3client

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func CreateBucket(ctx context.Context, s3Client *s3.Client) error {
	bucketName := "my-bucket-for-test-2"
	_, err := s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	var owned *types.BucketAlreadyOwnedByYou
	var exists *types.BucketAlreadyExists
	if err != nil && !errors.As(err, &owned) && !errors.As(err, &exists) {
		return fmt.Errorf("erro ao criar bucket: %w", err)
	}
	return nil
}

func ListBucket(ctx context.Context, s3Client *s3.Client) error {
	out, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("erro ao carregar as s3 buckets: %w", err)
	}

	fmt.Println("Conexão com Ministack S3 bem-sucedida - Buckets:")
	for _, bucket := range out.Buckets {
		fmt.Println("- ", aws.ToString(bucket.Name))
	}
	return nil
}
