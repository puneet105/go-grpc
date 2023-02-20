package main

import (
	"github.com/puneet105/go-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloServer struct{
	proto.DemoServiceServer
}

type EmployeeServer struct{
	proto.EmployeeServiceServer
	savedEmployee []*proto.EmployeeRequest
}

func main(){
	lisaddr, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalf("Failed to start server %v", err)
	}

	/*grpcServer := grpc.NewServer()
	proto.RegisterDemoServiceServer(grpcServer, &HelloServer{})
	log.Printf("Demo Service Server started at %v", lisaddr.Addr())
	if err := grpcServer.Serve(lisaddr); err != nil{
		log.Fatalf("Failed to start grpc server %v", err)
	}*/

	grpcServer := grpc.NewServer()
	proto.RegisterEmployeeServiceServer(grpcServer, &EmployeeServer{})
	log.Printf("Employee Service Server started at %v", lisaddr.Addr())
	if err := grpcServer.Serve(lisaddr); err != nil{
		log.Fatalf("Failed to start grpc server %v", err)
	}
}
