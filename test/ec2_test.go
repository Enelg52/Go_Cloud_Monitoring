package test

import (
	"Go_Cloud_Monitoring/pkg"
	"fmt"
	"testing"
)

func TestListEc2(t *testing.T) {
	inst, err := pkg.ListEc2("eu-north-1")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if inst != nil {
		fmt.Println("Test passed")
	}
}
