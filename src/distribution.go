package main

import (
	"context"
	"strconv"
	"time"

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
	callRefBase string = "awscloudfront-invalidator"
	quanitity   int32  = 1
)

func (dist Distribution) Invalidate(PathString string, client *cloudfront.Client) (string, error) {
	unqiueCallerRef := callRefBase + "-" + strconv.Itoa(int(time.Now().Unix()))

	invalidationBatch := &types.InvalidationBatch{
		CallerReference: &unqiueCallerRef,
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
