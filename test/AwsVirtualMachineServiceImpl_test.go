package test

import (
	"Go_Cloud_Monitoring/pkg"
	"testing"
	"time"
)

var (
	cloudProfileName string
	region           string
	machineToStart   int
)

func AwsVirtualMachineServiceImpl_BeforeAll() {
	cloudProfileName = "test"
	region = "eu-north-1"
	machineToStart = 0
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

func Test_AwsVirtualMachineServiceImpl_StartInstances(t *testing.T) {
	// given : Instancier un nouvel objet virtual machine
	awsVirtualMachineService, err := pkg.NewAwsVirtualMachineServiceImpl(cloudProfileName, region)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	i, err := awsVirtualMachineService.DescribeVirtualMachines()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// when : Describe virtual machine
	err = awsVirtualMachineService.StartInstances(i[machineToStart].GetInstanceId())
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// then verifier que ça soit les bonnes valeurs
	i, err = awsVirtualMachineService.DescribeVirtualMachines()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if i[machineToStart].GetStatus() == "stopped" {
		t.Errorf("Error: %s", err)
	}
}

func Test_AwsVirtualMachineServiceImpl_StopInstances(t *testing.T) {
	// given : Instancier un nouvel objet virtual machine
	awsVirtualMachineService, err := pkg.NewAwsVirtualMachineServiceImpl(cloudProfileName, region)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	i, err := awsVirtualMachineService.DescribeVirtualMachines()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// when : Describe virtual machine
	err = awsVirtualMachineService.StopInstances(i[machineToStart].GetInstanceId())
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// then verifier que ça soit les bonnes valeurs après une seconde
	time.Sleep(1 * time.Second)
	i, err = awsVirtualMachineService.DescribeVirtualMachines()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if i[machineToStart].GetStatus() == "running" {
		t.Errorf("Error: %s", err)
	}
}

func Test_AwsVirtualMachineServiceImpl_Exception_WrongRegion(t *testing.T) {
	// given : Instancier un nouvel objet virtual machine
	region = "eu-noorth-1" //region qui n'existe pas
	_, err := pkg.NewAwsVirtualMachineServiceImpl(cloudProfileName, region)
	if err == nil {
		t.Errorf("Error: %s", "No exeption")
	}
}
