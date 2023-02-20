package main

import (
	"context"
	"github.com/puneet105/go-grpc/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirStream(client proto.DemoServiceClient, names *proto.NameList){
	log.Println("Bidirectional streaming started....!!")
	stream, err := client.SayHelloBiDirectionnalStreaming(context.Background())
	if err != nil{
		log.Fatalf("couldn't send names : %v",err)
	}

	wait := make(chan struct{})

	go func() {
		for{
			message, err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil{
				log.Fatalf("error while streaming : %v",err)
			}
			log.Println(message)
		}
		close(wait)
	}()

	 for _, name := range names.Name{
		 req := &proto.HelloRequest{
			 Name: name,
		 }
		 if err := stream.Send(req); err != nil{
			 log.Fatalf("error while sending : %v", err)
		 }
		 time.Sleep(2*time.Second)
	 }
	 stream.CloseSend()
	 <-wait
	 log.Println("Bidirectional streaming finished....!!")
}
