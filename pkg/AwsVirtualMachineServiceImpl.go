package pkg

import (
	"Go_Cloud_Monitoring/internal"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AwsVirtualMachineServiceImpl struct {
	cloudProfileName string
	region           string
}

func NewAwsVirtualMachineServiceImpl(cloudProfileName string, region string) *AwsVirtualMachineServiceImpl {
	return &AwsVirtualMachineServiceImpl{cloudProfileName: cloudProfileName, region: region}
}

func (awsVmService *AwsVirtualMachineServiceImpl) GetCloudProfileName() string {
	return awsVmService.cloudProfileName
}

func (awsVmService *AwsVirtualMachineServiceImpl) GetRegion() string {
	return awsVmService.region
}

func (awsVmService *AwsVirtualMachineServiceImpl) SetCloudProfileName(cloudProfileName string) {
	awsVmService.cloudProfileName = cloudProfileName
}

func (awsVmService *AwsVirtualMachineServiceImpl) SetRegion(region string) {
	awsVmService.region = region
}

func (awsVmService *AwsVirtualMachineServiceImpl) DescribeVirtualMachines() ([]*internal.VirtualMachine, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsVmService.GetRegion())},
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
