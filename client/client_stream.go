package main

import (
	"context"
	"github.com/puneet105/go-grpc/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client proto.DemoServiceClient, names *proto.NameList){
	log.Println("Client Streaming Started....!!")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil{
		log.Fatalf("Couldn't send names : %v", err)
	}
	for _, name := range names.Name{
		req := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil{
			log.Fatalf("error while sending %v",err)
		}
		log.Println("sent the request with name : ",name)
		time.Sleep(2*time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Println("Client Streaming finished...!!")
	if err != nil{
		log.Fatalf("error while receiving %v",err)
	}
	log.Println(res.Message)
}
