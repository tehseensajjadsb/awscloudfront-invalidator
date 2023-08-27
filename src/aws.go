package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

var CloudfrontClient = cloudfront.NewFromConfig(GetAwsAuthConfig(Region))

func GetAwsAuthConfig(regionString string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(regionString))
	if err != nil {
		log.Fatalf("Unable to authenticate with AWS. %v", err)
	}
	return cfg
}

func GetCallerIdentity(cfg aws.Config) string {
	stsClient := sts.NewFromConfig(cfg)
	callerIdentityOutput, err := stsClient.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve current session details, %v", err)
	}

	return fmt.Sprintf("ACCOUNT_ID: %10v\nARN: %10v\nUSER_ID: %10v\n", *callerIdentityOutput.Account, *callerIdentityOutput.Arn, *callerIdentityOutput.UserId)
}
