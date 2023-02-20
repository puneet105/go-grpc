package main

import (
	"github.com/puneet105/go-grpc/proto"
	"io"
	"log"
)

func (s *HelloServer)SayHelloBiDirectionnalStreaming(stream proto.DemoService_SayHelloBiDirectionnalStreamingServer)error{
	for{
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil{
			return err
		}
		log.Println("got request with name : ", req.Name)
		res := &proto.HelloResponse{
			Message: req.Name,
		}
		if err := stream.Send(res); err != nil{
			return err
		}
	}
}
