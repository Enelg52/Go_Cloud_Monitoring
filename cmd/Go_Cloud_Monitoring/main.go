package main

import (
	"Go_Cloud_Monitoring/pkg"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	action := flag.String("a", "list", "Specify the action to execute")
	region := flag.String("r", "eu-north-1", "Region")
	flag.Parse()
	// Load session from shared config
	sess, err := session.NewSession(&aws.Config{
		// Specify region
		Region: aws.String(*region)},
	)
	if err != nil {
		pkg.ErrPrint(err)
	}

	// Create new EC2 client
	ec2Svc := ec2.New(sess)

	// Check action
	switch *action {
	case "list":
		inst, err := pkg.ListEc2(ec2Svc)
		if err != nil {
			pkg.ErrPrint(err)
		}
		fmt.Println(inst.String())
	}
}
