package s3client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Upload(ctx context.Context, s3Client *s3.Client, bucket, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	key := filepath.Base(filePath) //nome do objeto = nome do arquivo

	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("erro ao fazer o upload: %w", err)
	}
	fmt.Printf("upload concluído: %s -> %s/%s\n", filePath, bucket, key)
	return nil
}
