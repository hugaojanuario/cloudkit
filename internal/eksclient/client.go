package eksclient

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func NewClient(cfg aws.Config) *eks.Client {
	return eks.NewFromConfig(cfg, func(o *eks.Options) {
		o.BaseEndpoint = aws.String("http://localhost:4566")
	})
}
