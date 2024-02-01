package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)


func main(){

	session,err:=session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})
  if err!=nil{
	log.Fatal("failed to create a session",err)
	os.Exit(1)
  }

  serviceClient:=ec2.New(session)

  parameters:=&ec2.RunInstancesInput{
	ImageId: aws.String("ami-019ef0704f578a614"),
	InstanceType: aws.String("t2.micro"),
	MinCount: aws.Int64(1),
	MaxCount: aws.Int64(1),

  }
  createInstance,err:=serviceClient.RunInstances(parameters)

  if err!=nil {
	log.Fatal("Failed to create an instances",err)
	os.Exit(1)
	
  }

  instanceID:=*createInstance.Instances[0].InstanceId
  fmt.Println("created an ec2 instances and here it is ",instanceID)
}