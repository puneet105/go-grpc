package main

import (
	"github.com/puneet105/go-grpc/proto"
	"log"
	"time"
)

func (s *HelloServer)SayHelloServerStreaming(req *proto.NameList,  stream proto.DemoService_SayHelloServerStreamingServer)error{
	log.Println("Got request for names : ",req.Name)
	for _, nam := range req.Name{
		res := &proto.HelloResponse{
			Message: "Hello From =>" + nam,
		}
		if err := stream.Send(res); err != nil{
			return err
		}
		time.Sleep(2*time.Second)
	}
	return nil
}
