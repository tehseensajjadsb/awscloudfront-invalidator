package main

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

const CallerReferencePrefix = "awscloudfront-invalidator"

type DistributionInvalidator interface {
	Invalidate(string) string
	WaitForInvalidation()
}

type Distribution struct {
	Id string
}

func (dist *Distribution) Invalidate(invalidationPaths []string, client *cloudfront.Client) (string, error) {
	currentTime := int(time.Now().Unix())
	unqiueCallerRef := CallerReferencePrefix + "-" + strconv.Itoa(currentTime)

	pathsCount := int32(len(invalidationPaths))

	invalidationBatch := &types.InvalidationBatch{
		CallerReference: &unqiueCallerRef,
		Paths: &types.Paths{
			Quantity: &pathsCount,
			Items:    invalidationPaths,
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
