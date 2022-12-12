package test

import (
	"Go_Cloud_Monitoring/internal"
	"testing"
)

func TestProperties_NominalCase_GetAllProperties(t *testing.T) {
	name := "test"
	os := "linux"
	imageId := "1234"

	// given : Instancier un nouvel objet virtual machine

	virtualMachine := internal.NewVirtualMachine(name, os, imageId)

	// when : Get all properties
	virtualMachineName := virtualMachine.GetName()
	virtualMachineOs := virtualMachine.GetOs()
	virtualMachineImageId := virtualMachine.GetImageId()

	// then verifier que ça soit les bonnes valeurs
	if virtualMachineName != name {
		t.Errorf("Error: %s", "Name is not equal")
	}
	if virtualMachineOs != os {
		t.Errorf("Error: %s", "Os is not equal")
	}
	if virtualMachineImageId != imageId {
		t.Errorf("Error: %s", "ImageId is not equal")
	}
}
