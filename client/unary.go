package main

import (
	"context"
	"github.com/puneet105/go-grpc/proto"
	"log"
	"time"
)

func callSayHello( client proto.DemoServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &proto.NoParam{})
	if err != nil{
		log.Fatalf("couldn't get response from server %v",err)
	}
	log.Printf("Message is : %s", resp.Message)

}

