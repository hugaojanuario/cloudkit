package main

import (
	"context"

	"github.com/hugaojanuario/cloudkit/internal/s3client"
)

func main() {
	ctx := context.Background()

	client := s3client.NewClient(ctx)
	s3client.CreateBucket(client, ctx)
	s3client.ListBucket(client, ctx)

}
