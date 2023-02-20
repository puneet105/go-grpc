package main

import (
	"github.com/puneet105/go-grpc/proto"
	"io"
	"log"
)

func(s *HelloServer)SayHelloClientStreaming(stream proto.DemoService_SayHelloClientStreamingServer)error{
	var messages []string
	for{
		req, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&proto.MessageResponseList{Message: messages})
		}
		if err != nil{
			return err
		}
		log.Println("Got request for names : ",req.Name)
		messages = append(messages, "Hello From ",req.Name)
	}
}
