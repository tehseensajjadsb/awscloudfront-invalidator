package main

import (
	"fmt"
)

func main() {
	Init()

	awsAuthConfig := GetAwsAuthConfig(Region)
	fmt.Printf("%v", GetCallerIdentity(awsAuthConfig))
}
