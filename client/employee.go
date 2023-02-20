package main

import (
	"context"
	"github.com/puneet105/go-grpc/proto"
	"io"
	"log"
	"time"
)

func callCreateEmployee(client proto.EmployeeServiceClient, employee *proto.EmployeeRequest){
	resp, err := client.CreateEmployee(context.Background(), employee)
	if err != nil{
		log.Fatalf("Couldn't create employee : %v", err)
	}
	if resp.Result{
		log.Println("New employee has been added with id : ", resp.Id)
	}
}

func callGetEmployees(client proto.EmployeeServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := client.GetEmployee(ctx, &proto.NoParams{})
	if err != nil{
		log.Fatalf("error fetching employee details : %v",err)
	}
	for{
		emp, err := stream.Recv()
		//log.Printf("receiving employee detail :%+v", emp)
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("error is : %v",err)
		}
		log.Println("employee details are: ", emp)
	}
}