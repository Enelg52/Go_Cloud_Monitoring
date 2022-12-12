package test

import (
	"Go_Cloud_Monitoring/pkg"
	"testing"
)

var (
	cloudProfileName = "test"
	region           = "eu-north-1"
)

func Test_AwsVirtualMachineServiceImpl_Properties_NominalCase_GetAllProperties(t *testing.T) {
	// given : Instancier un nouvel objet virtual machine

	awsVirtualMachineService := pkg.NewAwsVirtualMachineServiceImpl(cloudProfileName, region)

	// when : Get all properties
	awsVirtualMachineServiceProfileName := awsVirtualMachineService.GetCloudProfileName()
	awsVirtualMachineServiceRegion := awsVirtualMachineService.GetRegion()

	// then verifier que ça soit les bonnes valeurs
	if awsVirtualMachineServiceProfileName != cloudProfileName {
		t.Errorf("Error: %s", "cloudProfileName is not equal")
	}
	if awsVirtualMachineServiceRegion != region {
		t.Errorf("Error: %s", "Region is not equal")
	}
}

func Test_AwsVirtualMachineServiceImpl_DescribeInstance(t *testing.T) {
	// given : Instancier un nouvel objet virtual machine
	awsVirtualMachineService := pkg.NewAwsVirtualMachineServiceImpl(cloudProfileName, region)
	i, err := awsVirtualMachineService.DescribeVirtualachines()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if i == nil {
		t.Errorf("Error: %s", "Instance is nil")
	}
}
