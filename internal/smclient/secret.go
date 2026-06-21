package smclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func CreateSecret(ctx context.Context, client *secretsmanager.Client, name, value string) (string, error) {
	output, err := client.CreateSecret(ctx, &secretsmanager.CreateSecretInput{
		Name:         aws.String(name),
		SecretString: aws.String(value),
	})
	if err != nil {
		return "", fmt.Errorf("erro ao criar a secret: %w", err)
	}

	return aws.ToString(output.ARN), nil
}

func GetSecret(ctx context.Context, client *secretsmanager.Client, name string) (string, error) {
	output, err := client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	})
	if err != nil {
		return "", fmt.Errorf("buscar secret: %w", err)
	}
	return aws.ToString(output.SecretString), nil
}

func ListSecrets(ctx context.Context, client *secretsmanager.Client) ([]string, error) {
	output, err := client.ListSecrets(ctx, &secretsmanager.ListSecretsInput{})
	if err != nil {
		return nil, fmt.Errorf("listar secrets: %w", err)
	}

	names := make([]string, 0, len(output.SecretList))
	for _, secret := range output.SecretList {
		names = append(names, aws.ToString(secret.Name))
	}
	return names, nil
}
