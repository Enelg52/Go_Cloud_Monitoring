package pkg

import "Go_Cloud_Monitoring/internal"

type IVirtualMachineService interface {
	DescribeVirtualMachines() ([]*internal.VirtualMachine, error)
	StartInstances(instanceId string) error
	StopInstances(instanceId string) error
}
