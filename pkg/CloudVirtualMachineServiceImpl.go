package pkg

import (
	"Go_Cloud_Monitoring/internal"
)

// Function de test qui simule un autre cloud provider pour tester le polymorphisme

type CloudVirtualMachineServiceImpl struct {
	cloudProfileName string
	region           string
}

func NewCloudVirtualMachineServiceImpl(cloudProfileName string, region string) *CloudVirtualMachineServiceImpl {
	return &CloudVirtualMachineServiceImpl{cloudProfileName: cloudProfileName, region: region}
}

func (CloudVmService *CloudVirtualMachineServiceImpl) GetCloudProfileName() string {
	return CloudVmService.cloudProfileName
}

func (CloudVmService *CloudVirtualMachineServiceImpl) GetRegion() string {
	return CloudVmService.region
}

func (CloudVmService *CloudVirtualMachineServiceImpl) SetCloudProfileName(cloudProfileName string) {
	CloudVmService.cloudProfileName = cloudProfileName
}

func (CloudVmService *CloudVirtualMachineServiceImpl) SetRegion(region string) {
	CloudVmService.region = region
}

func (CloudVmService *CloudVirtualMachineServiceImpl) DescribeVirtualMachines() ([]*internal.VirtualMachine, error) {
	vmList := make([]*internal.VirtualMachine, 1)
	vmList[0] = internal.NewVirtualMachine("je", "suis", "un test")
	return vmList, error(nil)
}
