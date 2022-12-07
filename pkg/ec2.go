package pkg

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

func ListEc2(ec2Svc *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
	// Call to get detailed information on each instance
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		return nil, err
	} else {
		return result, error(nil)
	}
}
