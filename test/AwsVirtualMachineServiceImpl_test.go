package test

import (
	"Go_Cloud_Monitoring/pkg"
	"testing"
)

var (
	cloudProfileName string
	region           string
)

func AwsVirtualMachineServiceImpl_BeforeAll() {
	cloudProfileName = "test"
	region = "eu-north-1"
}

func Test_AwsVirtualMachineServiceImpl_Properties_NominalCase_GetAllProperties(t *testing.T) {
	// given : Instancier un nouvel objet virtual machine
	awsVirtualMachineService, err := pkg.NewAwsVirtualMachineServiceImpl(cloudProfileName, region)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

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
	awsVirtualMachineService, err := pkg.NewAwsVirtualMachineServiceImpl(cloudProfileName, region)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// when : Describe virtual machine
	i, err := awsVirtualMachineService.DescribeVirtualMachines()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// then verifier que ça soit les bonnes valeurs
	if i == nil {
		t.Errorf("Error: %s", "Instance is nil")
	}
}

func Test_AwsVirtualMachineServiceImpl_Exeption_WrongRegion(t *testing.T) {
	// given : Instancier un nouvel objet virtual machine
	region = "eu-noorth-1"
	_, err := pkg.NewAwsVirtualMachineServiceImpl(cloudProfileName, region)
	if err == nil {
		t.Errorf("Error: %s", "the exeption does not work")
	}
}
