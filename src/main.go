package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var (
	Region                   string
	InputId                  string
	InputAlias               string
	InputOriginPath          string
	Paths                    string
	GetCurrentProfileDetails bool
)

var Distribution Invalidatable

func Init() {
	flag.StringVar(&Region, "region", "us-east-1", "AWS Region")
	flag.BoolVar(&GetCurrentProfileDetails, "whoami", false, "Get current profile details and exit, similar to output of 'aws sts get-caller-identity'")
	flag.StringVar(&InputId, "distribution-id", "", "Cloudfront distribution id")
	flag.StringVar(&InputAlias, "distribution-alias", "", "Cloudfront distribution configured alias domain")
	flag.StringVar(&InputOriginPath, "distribution-originpath", "", "Cloudfront distribution configured origin path")
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

	// TODO: Find a better way than these if statements

	if InputId != "" {
		Distribution = DistributionById{
			Id: InputId,
		}
	}

	if InputAlias != "" {
		Distribution = DistributionByAlias{
			Alias: InputAlias,
		}
	}

	if InputOriginPath != "" {
		Distribution = DistributionByOriginPath{
			OriginPath: InputOriginPath,
		}
	}

	if Paths == "" {
		log.Fatal("You must provide a path to invalidate")
	}

	invalidationPaths := strings.Split(Paths, ",")

	invalidationId, err := Invalidate(invalidationPaths, Distribution)
	if err != nil {
		log.Fatalf("Failed to invalidate distribution %v\n", err)
	}

	fmt.Printf("%s", invalidationId)
}
