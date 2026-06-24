package eksclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

func CreateCluster(ctx context.Context, client *eks.Client, nameCluster, roleArn string, subnetIds []string) (string, error) {
	output, err := client.CreateCluster(ctx, &eks.CreateClusterInput{
		Name:    aws.String(nameCluster),
		RoleArn: aws.String(roleArn),
		ResourcesVpcConfig: &types.VpcConfigRequest{
			SubnetIds: subnetIds,
		},
	})
	if err != nil {
		return "", fmt.Errorf("erro ao criar o cluster: %w", err)
	}
	return aws.ToString(output.Cluster.Arn), nil
}

func DescribeCluster(ctx context.Context, client *eks.Client, nameCluster string) (string, error) {
	output, err := client.DescribeCluster(ctx, &eks.DescribeClusterInput{
		Name: aws.String(nameCluster),
	})
	if err != nil {
		return "", fmt.Errorf("erro ao buscar o cluster: %w", err)
	}
	return aws.ToString(output.Cluster.Arn), nil
}

func ListCluster(ctx context.Context, client *eks.Client) ([]string, error) {
	output, err := client.ListClusters(ctx, &eks.ListClustersInput{})
	if err != nil {
		return nil, fmt.Errorf("erro ao listar os cluters: %w", err)
	}
	return output.Clusters, nil
}
