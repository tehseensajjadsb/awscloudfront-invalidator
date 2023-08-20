package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func main() {
	// Parse commandline flags
	Init()
	awsAuthConfig := GetAwsAuthConfig(Region)

	if GetCurrentProfileDetails {
		callerIdentityOutput := GetCallerIdentity(awsAuthConfig)
		fmt.Print(callerIdentityOutput)
		return
	}

	if CloudfrontDistributionId != "" {
		cfnClient := cloudfront.NewFromConfig(awsAuthConfig)

		distribution := Distribution{
			Id: CloudfrontDistributionId,
		}

		if InvalidationPath == "" {
			log.Fatal("You must provide a path to invalidate")
		}

		invalidationId, err := distribution.Invalidate(InvalidationPath, cfnClient)
		if err != nil {
			log.Fatalf("Failed to invalidation distribution %v\n%v", CloudfrontDistributionId, err)
		}

		fmt.Printf("%s", invalidationId)
	}
}
