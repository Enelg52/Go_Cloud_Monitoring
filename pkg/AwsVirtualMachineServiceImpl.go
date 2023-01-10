package pkg

import (
	"Go_Cloud_Monitoring/internal"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AwsVirtualMachineServiceImpl struct {
	cloudProfileName string
	region           string
}

func NewAwsVirtualMachineServiceImpl(cloudProfileName string, region string) (*AwsVirtualMachineServiceImpl, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}
	ec2Svc := ec2.New(sess)
	//Test the creds
	_, err = ec2Svc.DescribeInstances(nil)

	if err != nil {
		err = errors.New("newAws_badRegion_throwException (invalid region)")
	}
	return &AwsVirtualMachineServiceImpl{cloudProfileName: cloudProfileName, region: region}, err
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

		//Get size off the 2d array
		size := 0
		for _, inst := range result.Reservations {
			for range inst.Instances {
				size++
			}
		}
		//make the return array
		vmList := make([]*internal.VirtualMachine, size)

		id := 0
		for _, inst := range result.Reservations {
			for _, val := range inst.Instances {
				vmList[id] = internal.NewVirtualMachine(
					*val.Tags[0].Value,              //name
					*val.PlatformDetails,            //os
					*val.ImageId,                    //imageId
					*val.InstanceId,                 //instanceId
					*val.InstanceType,               //instanceType
					internal.State(*val.State.Name), //state
				)
				id++
			}
		}
		return vmList, error(nil)
	}
}

func (awsVmService *AwsVirtualMachineServiceImpl) StartInstances(instanceId string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsVmService.GetRegion())},
	)
	if err != nil {
		return err
	}
	ec2Svc := ec2.New(sess)
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
	}
	_, err = ec2Svc.StartInstances(input)
	if err != nil {
		return err
	}
	return nil
}

func (awsVmService *AwsVirtualMachineServiceImpl) StopInstances(instanceId string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsVmService.GetRegion())},
	)
	if err != nil {
		return err
	}

	ec2Svc := ec2.New(sess)
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
	}
	_, err = ec2Svc.StopInstances(input)
	if err != nil {
		return err
	}
	return nil
}
