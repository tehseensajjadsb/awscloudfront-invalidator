package main

import "flag"

var (
	Region                   string
	CloudfrontDistributionId string
	GetCurrentProfileDetails bool
)

func Init() {
	flag.StringVar(&Region, "region", "us-east-1", "AWS Region")

	flag.BoolVar(&GetCurrentProfileDetails, "whoami", false, "Get current profile details, similar to output of 'aws sts get-caller-identity'")

	flag.StringVar(&CloudfrontDistributionId, "distribution-id", "", "Cloudfront distribution id")
	flag.Parse()
}
