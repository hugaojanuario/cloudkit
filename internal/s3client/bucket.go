package s3client

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func CreateBucket(s3Client *s3.Client, ctx context.Context) {
	bucketName := "my-bucket-for-test"
	_, err := s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	var owned *types.BucketAlreadyOwnedByYou
	var exists *types.BucketAlreadyExists
	if err != nil && !errors.As(err, &owned) && !errors.As(err, &exists) {
		log.Fatalf("erro ao criar bucket: %v", err)
	}

}

func ListBucket(s3Client *s3.Client, ctx context.Context) {
	out, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("erro ao carregar as s3 buckets: %v", err)
	}

	fmt.Println("Conexão com Ministack S3 bem-sucedida - Buckets:")
	for _, bucket := range out.Buckets {
		fmt.Println("- ", aws.ToString(bucket.Name))
	}
}
