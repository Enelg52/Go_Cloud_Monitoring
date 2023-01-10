package test

import (
	"Go_Cloud_Monitoring/internal"
	"testing"
)

var (
	name           string
	osName         string
	imageId        string
	instanceId     string
	instanceType   string
	status         internal.State
	virtualMachine *internal.VirtualMachine
)

func VirtalMachine_BeforeAll() {
	name = "test"
	osName = "test"
	imageId = "123423"
	instanceId = "123123123"
	status = "running"
}

// Before each
func VirtualMachine_BeforeEach() {
	virtualMachine = internal.NewVirtualMachine(name, osName, imageId, instanceId, instanceType, status)
}

func Test_VirtualMachine_Properties_NominalCase_GetAllProperties(t *testing.T) {

	// given : Instancier un nouvel objet virtual machine
	VirtualMachine_BeforeEach()
	// when : Get all properties
	virtualMachineName := virtualMachine.GetName()
	virtualMachineOs := virtualMachine.GetOs()
	virtualMachineImageId := virtualMachine.GetImageId()
	virtualMachineInstanceId := virtualMachine.GetInstanceId()
	virtualMachineInstanceType := virtualMachine.GetInstanceType()
	virtualMachineStatus := virtualMachine.GetStatus()

	// then verifier que ça soit les bonnes valeurs
	if virtualMachineName != name {
		t.Errorf("Error: %s", "Name is not equal")
	}
	if virtualMachineOs != osName {
		t.Errorf("Error: %s", "Os is not equal")
	}
	if virtualMachineImageId != imageId {
		t.Errorf("Error: %s", "ImageId is not equal")
	}
	if virtualMachineInstanceId != instanceId {
		t.Errorf("Error: %s", "Instance id is not equal")
	}
	if virtualMachineInstanceType != instanceType {
		t.Errorf("Error: %s", "Instance type is not equal")
	}
	if virtualMachineStatus != status {
		t.Errorf("Error: %s", "Status is not equal")
	}
}
