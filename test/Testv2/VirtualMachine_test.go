package test

import (
	"Go_Cloud_Monitoring/internal"
	"os"
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
Dans cette variante, on utilise la méthode TestMain() qui à été faite pour des tests unitairs. Dans ce main, on peut définir la séquence de nos test. Les before each/ after each sont simulé dans notre méthode de test.
Problème : Il peut y avoir qu'un TestMain par package. Je vois donc plusieurs solutions :
- Faire un package par objet a tester
- Faire un before all / after all pour tout les tests.
==============================================================
*/

// Main de TOUTE les methodes de tests
func TestMain(m *testing.M) {
	beforeAll()
	exitCode := m.Run()
	afterAll()
	os.Exit(exitCode)
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

func afterEach() {

}

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
