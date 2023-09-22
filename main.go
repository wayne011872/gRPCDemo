package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pbh "github.com/wayne011872/gRPCDemo/proto/hello"
	"github.com/wayne011872/gRPCDemo/service"
)

const (
	port = ":30000"
)
func main() {
	// Create gRPC Server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("gRPC server is running.")
	pbh.RegisterGreeterServer(s, &service.MessageService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}