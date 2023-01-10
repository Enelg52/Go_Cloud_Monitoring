package internal

type VirtualMachine struct {
	name         string
	os           string
	imageId      string
	instanceId   string
	status       State
	instanceType string
}

type State string

const (
	Running      State = "Running"
	Pending      State = "Pending"
	ShuttingDown State = "ShuttingDown"
	Stopping     State = "Stopping"
	Stopped      State = "Stopped"
	Terminated   State = "Terminated"
)

func NewVirtualMachine(name, os, imageId, instanceId, instanceType string, status State) *VirtualMachine {
	return &VirtualMachine{name: name, os: os, imageId: imageId, status: status, instanceId: instanceId, instanceType: instanceType}
}

func (vm *VirtualMachine) GetName() string {
	return vm.name
}

func (vm *VirtualMachine) GetOs() string {
	return vm.os
}

func (vm *VirtualMachine) GetImageId() string {
	return vm.imageId
}

func (vm *VirtualMachine) GetStatus() State {
	return vm.status
}

func (vm *VirtualMachine) GetInstanceId() string {
	return vm.instanceId
}

func (vm *VirtualMachine) GetInstanceType() string {
	return vm.instanceType
}
