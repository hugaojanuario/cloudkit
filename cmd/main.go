package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hugaojanuario/cloudkit/internal/awsconfig"
	"github.com/hugaojanuario/cloudkit/internal/s3client"
	"github.com/hugaojanuario/cloudkit/internal/sqsclient"
)

func main() {
	//Setup AWS
	ctx := context.Background()
	cfg, err := awsconfig.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//S3
	s3c := s3client.NewClient(cfg)

	bucket := "my-bucket-for-test"
	filePath := "teste.txt"

	if err := s3client.CreateBucket(ctx, s3c); err != nil {
		log.Fatal(err)
	}
	if err := s3client.ListBuckets(ctx, s3c); err != nil {
		log.Fatal(err)
	}

	if err := s3client.Upload(ctx, s3c, bucket, filePath); err != nil {
		log.Fatal(err)
	}

	if err := s3client.ListObjects(ctx, s3c, bucket); err != nil {
		log.Fatal(err)
	}

	if err := s3client.GetObject(ctx, s3c, bucket, filePath); err != nil {
		log.Fatal(err)
	}

	//SQS
	sqsc := sqsclient.NewClient(cfg)

	queueName := "my-queue-for-test"

	queueURL, err := sqsclient.CreateQueue(ctx, sqsc, queueName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fila criada:", queueURL)

	message, err := sqsclient.SendMessage(ctx, sqsc, queueURL, "Hello World")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("messagem enviada com sucesso - MessageId: ", message)

	receive, err := sqsclient.ReceiveMessage(ctx, sqsc, queueURL)
	if err != nil {
		log.Fatal(err)
	}
	if len(receive) == 0 {
		fmt.Println("nenhuma mensagem na fila")
	}

	for _, msg := range receive {
		fmt.Println("body:", aws.ToString(msg.Body))
		fmt.Println("receipt handle:", aws.ToString(msg.ReceiptHandle))

		if err := sqsclient.DeleteMessage(ctx, sqsc, queueURL, aws.ToString(msg.ReceiptHandle)); err != nil {
			log.Println("erro ao deletar:", err)
			continue
		}
		fmt.Println("mensagem deletada com sucesso")
	}

}
