package service

import(
	"context"
	pbh "github.com/wayne011872/gRPCDemo/proto/hello"
)

type MessageService struct{
	pbh.UnimplementedGreeterServer
}

func (m MessageService) SayHello(ctx context.Context, in *pbh.HelloRequest) (*pbh.HelloReply, error) {
	return &pbh.HelloReply{Message: "Server say hello to " + in.GetName()}, nil
}