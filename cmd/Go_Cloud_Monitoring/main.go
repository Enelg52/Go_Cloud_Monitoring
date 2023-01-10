package main

import (
	"Go_Cloud_Monitoring/internal"
	"Go_Cloud_Monitoring/pkg"
	"fmt"
)

func main() {
	vm1, err := pkg.NewAwsVirtualMachineServiceImpl("", "eu-north-1")
	if err != nil {
		pkg.ErrPrint(err)
		return
	}
	//err = startVirtualMachine(vm1, "i-03ce507da047dd4f7")
	//err = stopVirtualMachine(vm1, "i-03ce507da047dd4f7")
	if err != nil {
		pkg.ErrPrint(err)
	}
	decr1, err := describeVirtualMachines(vm1)
	if err != nil {
		pkg.ErrPrint(err)
		return
	}
	for _, v := range decr1 {
		fmt.Println(v)
	}
}

func describeVirtualMachines(i pkg.IVirtualMachineService) ([]*internal.VirtualMachine, error) {
	machines, err := i.DescribeVirtualMachines()
	if err != nil {
		return nil, err
	}
	return machines, error(nil)
}

func startVirtualMachine(i pkg.IVirtualMachineService, instanceId string) error {
	err := i.StartInstances(instanceId)
	if err != nil {
		return err
	}
	return nil
}

func stopVirtualMachine(i pkg.IVirtualMachineService, instanceId string) error {
	err := i.StopInstances(instanceId)
	if err != nil {
		return err
	}
	return nil
}
