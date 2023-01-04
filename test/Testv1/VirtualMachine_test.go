package test

import (
	"Go_Cloud_Monitoring/internal"
	"testing"
)

var (
	name           string
	opSys          string
	imageId        string
	virtualMachine *internal.VirtualMachine
)

/*
==============================================================
Dans cette variante, on utilise la méthode init() qui va être exécuté en premier. Ensuite, on va simuler le before each et le after each. Je n'ai pas trouvé comment simuler le after all.
Cette varainte devrait fonctionner pour nos besoins.
==============================================================
*/

// init est executé avant tout autre méthode.
func init() {
	beforeAll()
}

func beforeAll() {
	name = "test"
	opSys = "test"
	imageId = "123423"
}

// Before each
func beforeEach() {
	virtualMachine = internal.NewVirtualMachine(name, opSys, imageId)
}

// After each
func afterEach() {

}

// After all -> Impossible avec cette methode
func afterAll() {

}

func Test_VirtualMachine_Properties_NominalCase_GetAllProperties(t *testing.T) {

	// given : Instancier un nouvel objet virtual machine
	beforeEach()
	// S'execute quand la fonction se termine
	defer afterEach()
	// when : Get all properties
	virtualMachineName := virtualMachine.GetName()
	virtualMachineOs := virtualMachine.GetOs()
	virtualMachineImageId := virtualMachine.GetImageId()

	// then verifier que ça soit les bonnes valeurs
	if virtualMachineName != name {
		t.Errorf("Error: %s", "Name is not equal")
	}
	if virtualMachineOs != opSys {
		t.Errorf("Error: %s", "Os is not equal")
	}
	if virtualMachineImageId != imageId {
		t.Errorf("Error: %s", "ImageId is not equal")
	}
}
