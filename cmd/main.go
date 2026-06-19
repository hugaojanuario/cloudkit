package main

import (
	"context"
	"log"

	"github.com/hugaojanuario/cloudkit/internal/s3client"
)

func main() {
	ctx := context.Background()

	client, err := s3client.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	s3client.CreateBucket(ctx, client)
	s3client.ListBucket(ctx, client)

	bucket := "my-bucket-for-test-2"
	filePath := "C:\\Users\\hugo.santariosi\\GolandProjects\\cloudkit\\teste.txt"
	if err := s3client.Upload(ctx, client, bucket, filePath); err != nil {
		log.Fatal(err)
	}

}
