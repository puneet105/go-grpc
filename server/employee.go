package main

import (
	"context"
	"github.com/puneet105/go-grpc/proto"
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

func (s *EmployeeServer)GetEmployee(req *proto.NoParams, stream proto.EmployeeService_GetEmployeeServer)error{
	for _, emp := range s.savedEmployee{
		//log.Println("sending employee detail : ",emp)
		if err := stream.Send(emp); err != nil{
			return err
		}
	}
	return nil
}