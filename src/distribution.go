package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

const CallerReferencePrefix = "awscloudfront-invalidator"

type Invalidatable interface {
	GetDistributionId() (*string, error)
}

type DistributionById struct {
	Id string
}

func (dist DistributionById) GetDistributionId() (*string, error) {
	return &dist.Id, nil
}

type DistributionByAlias struct {
	Alias string
}

func (dist DistributionByAlias) GetDistributionId() (*string, error) {
	failedResponse := ""

	listMarker := "xyz"
	var maxItems int32 = 100
	params := cloudfront.ListDistributionsInput{
		Marker:   &listMarker,
		MaxItems: &maxItems,
	}
	distributions, err := CloudfrontClient.ListDistributions(context.TODO(), &params)
	if err != nil {
		return &failedResponse, nil
	}

	for _, distSummary := range distributions.DistributionList.Items {
		for _, alias := range distSummary.Aliases.Items {
			if alias == dist.Alias {
				return distSummary.Id, nil
			}
		}
	}

	return &failedResponse, errors.New(fmt.Sprintf("Could not find Cloudfront distribution with Alias: %s", dist.Alias))
}

type DistributionByOriginPath struct {
	OriginPath string
}

func (dist DistributionByOriginPath) GetDistributionId() (*string, error) {
	failedResponse := ""
	listMarker := "xyz"
	var maxItems int32 = 100
	params := cloudfront.ListDistributionsInput{
		Marker:   &listMarker,
		MaxItems: &maxItems,
	}
	distributions, err := CloudfrontClient.ListDistributions(context.TODO(), &params)
	if err != nil {
		return &failedResponse, err
	}

	for _, distSummary := range distributions.DistributionList.Items {
		for _, origin := range distSummary.Origins.Items {
			if dist.OriginPath == *origin.OriginPath {
				return distSummary.Id, nil
			}
		}
	}

	return &failedResponse, errors.New(fmt.Sprintf("Could not find Cloudfront distribution with Origin Path: %s", dist.OriginPath))
}
