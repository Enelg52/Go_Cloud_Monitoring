package test

import (
	"os"
	"testing"
)

// Main de TOUTE les methodes de tests
func TestMain(m *testing.M) {
	VirtalMachine_BeforeAll()
	AwsVirtualMachineServiceImpl_BeforeAll()
	exitCode := m.Run()
	os.Exit(exitCode)
}
