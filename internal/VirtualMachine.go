package internal

type VirtualMachine struct {
	name    string
	os      string
	imageId string
}

func NewVirtualMachine(name, os, imageId string) *VirtualMachine {
	return &VirtualMachine{name: name, os: os, imageId: imageId}
}

func (vm *VirtualMachine) SetName(name string) {
	vm.name = name
}

func (vm *VirtualMachine) SetOs(os string) {
	vm.os = os
}

func (vm *VirtualMachine) SetImageId(imageId string) {
	vm.imageId = imageId
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
