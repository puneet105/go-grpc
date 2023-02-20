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

func callCreateMultipleEmployee(client proto.EmployeeServiceClient, employees *[]proto.EmployeeRequest){
	stream, err := client.CreateMultipleEmployee(context.Background())
	if err != nil{
		log.Fatalf("Couldn't create employee : %v", err)
	}
	wait := make(chan struct{})

	go func() {
		for{
			emp, err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil{
				log.Fatalf("error while streaming : %v",err)
			}
			log.Printf("Employee %+v",emp)
		}
		close(wait)
	}()

	for _,emp := range *employees{
		req := &proto.EmployeeRequest{
			Id: emp.Id,
			Name: emp.Name,
			Email: emp.Email,
			Phone: emp.Phone,
		}
		if err := stream.Send(req); err != nil{
			log.Fatalf("error while sending : %v", err)
		}
		time.Sleep(2*time.Second)
	}

	stream.CloseSend()
	<-wait
}

func callGetEmployees(client proto.EmployeeServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	employeeData := []*proto.EmployeeRequest{}
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
		employeeData = append(employeeData,emp)
	}

	log.Printf("employee details are: %+v", employeeData)
}