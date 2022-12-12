package main

import (
	"Go_Cloud_Monitoring/pkg"
	"fmt"
)

func main() {

	vm := pkg.NewAwsVirtualMachineServiceImpl("", "eu-north-1")
	a, err := vm.DescribeVirtualachines()
	if err != nil {
		pkg.ErrPrint(err)
	}
	//print all a
	for _, v := range a {
		fmt.Println(v)
	}

}
