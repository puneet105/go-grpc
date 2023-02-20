package main

import (
	"context"
	"github.com/puneet105/go-grpc/proto"
	"io"
	"log"
)

func (s *EmployeeServer)CreateEmployee(ctx context.Context, req *proto.EmployeeRequest)(*proto.EmployeeResponse, error){
	log.Println("Got a request to create employee : ", req)
	s.savedEmployee = append(s.savedEmployee, req)
	return &proto.EmployeeResponse{
		Id: req.Id,
		Result: true,
	},nil
}

func (s *EmployeeServer)CreateMultipleEmployee(stream proto.EmployeeService_CreateMultipleEmployeeServer)error{
	for{
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil{
			return err
		}
		log.Printf("got request for employee : %+v", req)
		s.savedEmployee = append(s.savedEmployee, req)
		res := &proto.EmployeeResponse{
			Id: req.Id,
			Result: true,
		}
		if err := stream.Send(res); err != nil{
			return err
		}
	}
}

func (s *EmployeeServer)GetEmployee(req *proto.NoParams, stream proto.EmployeeService_GetEmployeeServer)error{
	for _, emp := range s.savedEmployee{
		//log.Println("sending employee detail : ",emp)
		if err := stream.Send(emp); err != nil{
			return err
		}
	}
	return nil
}