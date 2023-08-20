package main

import "fmt"

func main() {
	awsAuthConfig := GetAwsAuthConfig("us-east-1")
	fmt.Printf("%v", GetCallerIdentity(awsAuthConfig))
}
