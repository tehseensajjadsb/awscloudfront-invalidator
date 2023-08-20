package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

type DistributionInvalidator interface {
	Invalidate(string) string
	WaitForInvalidation()
}

type Distribution struct {
	Id string
}

var (
	callerReference string = "awscloudfront-invalidator"
	quanitity       int32  = 1
)

func (dist Distribution) Invalidate(PathString string, client *cloudfront.Client) (string, error) {
	invalidationBatch := &types.InvalidationBatch{
		CallerReference: &callerReference,
		Paths: &types.Paths{
			Quantity: &quanitity,
			Items:    []string{PathString},
		},
	}
	invalidationInput := cloudfront.CreateInvalidationInput{
		DistributionId:    &dist.Id,
		InvalidationBatch: invalidationBatch,
	}

	invalidationResult, err := client.CreateInvalidation(context.TODO(), &invalidationInput)
	if err != nil {
		return "", err
	}

	invalidation := invalidationResult.Invalidation
	invalidationId := *invalidation.Id
	return invalidationId, nil
}
