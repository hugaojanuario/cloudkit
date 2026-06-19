package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("mock-id", "mock-key", "")),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalf("erro ao carregar as configs da AWS: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String("http://localhost:4566")
		o.UsePathStyle = true
	})

	//criando um bucket s3
	bucketName := "my-bucket-for-test"
	_, err = s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	var owned *types.BucketAlreadyOwnedByYou
	var exists *types.BucketAlreadyExists
	if err != nil && !errors.As(err, &owned) && !errors.As(err, &exists) {
		log.Fatalf("erro ao criar bucket: %v", err)
	}

	out, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("erro ao carregar as s3 buckets: %v", err)
	}

	fmt.Println("Conexão com Ministack S3 bem-sucedida - Buckets:")
	for _, bucket := range out.Buckets {
		fmt.Println("- ", aws.ToString(bucket.Name))
	}

}
