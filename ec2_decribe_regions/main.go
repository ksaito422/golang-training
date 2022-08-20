package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2DescribeRegionsAPI interface {
	DescribeRegions(ctx context.Context,
		params *ec2.DescribeRegionsInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error)
}

// AWSがサポートしているRegionを取得する [aws ec2 describe-regions]
func GetAllRegions(c context.Context, api EC2DescribeRegionsAPI, input *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	return api.DescribeRegions(c, input)
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	ec2Client := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeRegionsInput{}
	resp, err := GetAllRegions(context.TODO(), ec2Client, input)
	if err != nil {
		fmt.Println("Got an error retrieving regions")
		fmt.Println(err)
		os.Exit(1)
	}

	for _, region := range (*resp).Regions {
		fmt.Printf("%s %s\n", *region.RegionName, *region.Endpoint)
	}
}
