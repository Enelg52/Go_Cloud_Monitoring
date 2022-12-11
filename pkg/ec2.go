package pkg

import (
	"Go_Cloud_Monitoring/internal"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func ListEc2(region string) ([]*internal.VirtualMachine, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}
	ec2Svc := ec2.New(sess)
	// Call to get detailed information on each instance
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		return nil, err
	} else {
		// Create a new VirtualMachine
		vmList := make([]*internal.VirtualMachine, len(result.Reservations))
		for id, inst := range result.Reservations {
			vmList[id] = internal.NewVirtualMachine(*inst.Instances[0].InstanceId, *inst.Instances[0].ImageId, *inst.Instances[0].InstanceType)
		}
		return vmList, error(nil)
	}
}
