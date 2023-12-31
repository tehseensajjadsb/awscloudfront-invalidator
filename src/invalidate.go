package main

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

func Invalidate(invalidationPaths []string, distribution Invalidatable) (string, error) {
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

	distributionId, err := distribution.GetDistributionId()
	if err != nil {
		return "", err
	}

	invalidationInput := cloudfront.CreateInvalidationInput{
		DistributionId:    distributionId,
		InvalidationBatch: invalidationBatch,
	}

	invalidationResult, err := CloudfrontClient.CreateInvalidation(context.TODO(), &invalidationInput)
	if err != nil {
		return "", err
	}

	invalidation := invalidationResult.Invalidation
	invalidationId := *invalidation.Id
	return invalidationId, nil
}
