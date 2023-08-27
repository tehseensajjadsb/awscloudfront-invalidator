package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

var (
	Region                   string
	CloudfrontDistributionId string
	Paths                    string
	GetCurrentProfileDetails bool
)

func Init() {
	flag.StringVar(&Region, "region", "us-east-1", "AWS Region")
	flag.BoolVar(&GetCurrentProfileDetails, "whoami", false, "Get current profile details and exit, similar to output of 'aws sts get-caller-identity'")
	flag.StringVar(&CloudfrontDistributionId, "distribution-id", "", "Cloudfront distribution id")
	flag.StringVar(&Paths, "paths", "", "Path(s) to invalidate")
	flag.Parse()
}

func main() {
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

		if Paths == "" {
			log.Fatal("You must provide a path to invalidate")
		}

		invalidationPaths := strings.Split(Paths, ",")

		invalidationId, err := distribution.Invalidate(invalidationPaths, cfnClient)
		if err != nil {
			log.Fatalf("Failed to invalidation distribution %v\n%v", CloudfrontDistributionId, err)
		}

		fmt.Printf("%s", invalidationId)
	}
}
