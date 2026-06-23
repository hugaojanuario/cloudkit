package iamclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func CreateRole(ctx context.Context, client *iam.Client, roleName string, assumeRolePolicy string) (string, error) {
	output, err := client.CreateRole(ctx, &iam.CreateRoleInput{
		RoleName:                 aws.String(roleName),
		AssumeRolePolicyDocument: aws.String(assumeRolePolicy),
	})
	if err != nil {
		return "", fmt.Errorf("erro ao criar a role: %w", err)
	}

	return aws.ToString(output.Role.Arn), nil
}

func GetRole(ctx context.Context, client *iam.Client, roleName string) (string, error) {
	output, err := client.GetRole(ctx, &iam.GetRoleInput{
		RoleName: aws.String(roleName),
	})
	if err != nil {
		return "", fmt.Errorf("erro ao buscar a role: %w", err)
	}

	return aws.ToString(output.Role.Arn), nil
}

func ListRole(ctx context.Context, client *iam.Client) ([]string, error) {
	output, err := client.ListRoles(ctx, &iam.ListRolesInput{})
	if err != nil {
		return nil, fmt.Errorf("erro ao listar as roles: %w", err)
	}

	names := make([]string, 0, len(output.Roles))
	for _, r := range output.Roles {
		names = append(names, aws.ToString(r.RoleName))
	}
	return names, nil
}

func DeleteRole(ctx context.Context, client *iam.Client, roleName string) error {
	_, err := client.DeleteRole(ctx, &iam.DeleteRoleInput{
		RoleName: aws.String(roleName),
	})
	if err != nil {
		return fmt.Errorf("erro ao deletar a role: %w", err)
	}
	return nil
}
