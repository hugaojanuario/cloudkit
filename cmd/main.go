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

}
