package main

import (
	"Go_Cloud_Monitoring/pkg"
	"fmt"
)

func main() {
	a, err := pkg.ListEc2("eu-north-1")
	if err != nil {
		pkg.ErrPrint(err)
	}
	//print all a
	for _, v := range a {
		fmt.Println(v)
	}

}
