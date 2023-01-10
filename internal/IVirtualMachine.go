package internal

type IVirtualMachineService interface {
	NewVirtualMachine(name, os, imageId string) *VirtualMachine
	GetName() string
	GetOs() string
	GetImageId() string
	GetInstanceId() string
	GetStatus() State
	GetInstanceType() string
}
