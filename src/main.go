package main

import (
	"flag"
	"fmt"
	"os"
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
	flag.StringVar(&InputId, "id", "", "Cloudfront distribution id")
	flag.StringVar(&InputAlias, "alias", "", "Cloudfront distribution configured alias domain")
	flag.StringVar(&InputOriginPath, "origin-path", "", "Cloudfront distribution configured origin path")
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
		fmt.Fprint(os.Stderr, "You must provide a path to invalidate")
		os.Exit(1)
	}

	invalidationPaths := []string{Paths}
	if strings.Contains(Paths, ",") {
		invalidationPaths = strings.Split(strings.TrimSpace(Paths), ",")
		if len(invalidationPaths) == 0 {
			fmt.Fprintf(os.Stderr, "Incorrect invalidation Paths provided: %s", Paths)
			os.Exit(1)
		}
	}

	invalidationId, err := Invalidate(invalidationPaths, Distribution)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to invalidate distribution %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", invalidationId)
}
