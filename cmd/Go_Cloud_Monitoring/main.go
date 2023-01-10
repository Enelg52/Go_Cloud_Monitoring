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
