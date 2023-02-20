package main

import (
	"context"
	"github.com/puneet105/go-grpc/proto"
)

func (s *HelloServer)SayHello(ctx context.Context, req *proto.NoParam)(*proto.HelloResponse, error){
	return &proto.HelloResponse{
		Message: "Hello",
	},nil
}

