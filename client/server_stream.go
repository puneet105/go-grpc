package main

import (
	"context"
	"github.com/puneet105/go-grpc/proto"
	"io"
	"log"
)

func callSayHelloStream(client proto.DemoServiceClient, names *proto.NameList){
	log.Println("Streaming started ....!!")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil{
		log.Fatalf("couldn't send names : %v", err)
	}
	for{
		message, err := stream.Recv()
		if err == io.EOF{
			return
		}
		if err != nil{
			log.Fatalf("error while streaming %v",err)
		}
		log.Println(message)
	}
	log.Printf("Streaming Finished.")
}
