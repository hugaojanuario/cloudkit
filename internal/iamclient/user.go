package iamclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func CreateUser(ctx context.Context, client *iam.Client, username string) (string, error) {
	output, err := client.CreateUser(ctx, &iam.CreateUserInput{
		UserName: aws.String(username),
	})
	if err != nil {
		return "", fmt.Errorf("erro ao criar o usuario: %w", err)
	}

	return aws.ToString(output.User.Arn), nil
}

func GetUser(ctx context.Context, client *iam.Client, username string) (string, error) {
	output, err := client.GetUser(ctx, &iam.GetUserInput{
		UserName: aws.String(username),
	})
	if err != nil {
		return "", fmt.Errorf("erro ao encontrar o usuario: %w", err)
	}

	return aws.ToString(output.User.UserId), nil
}

func ListUser(ctx context.Context, client *iam.Client) ([]string, error) {
	output, err := client.ListUsers(ctx, &iam.ListUsersInput{})
	if err != nil {
		return nil, fmt.Errorf("erro ao listar os usuarios: %w", err)
	}

	usernames := make([]string, 0, len(output.Users))
	for _, iam := range output.Users {
		usernames = append(usernames, aws.ToString(iam.UserName))
	}
	return usernames, nil
}

func DeleteUser(ctx context.Context, client *iam.Client, username string) error {
	_, err := client.DeleteUser(ctx, &iam.DeleteUserInput{
		UserName: aws.String(username),
	})
	if err != nil {
		return fmt.Errorf("deletar usuário: %w", err)
	}
	return nil
}
