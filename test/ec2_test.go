package test

import (
	"Go_Cloud_Monitoring/pkg"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"testing"
)

func TestListEc2(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-north-1")},
	)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	ec2Svc := ec2.New(sess)
	inst, err := pkg.ListEc2(ec2Svc)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if inst != nil {
		fmt.Println("Test passed")
	}
}
