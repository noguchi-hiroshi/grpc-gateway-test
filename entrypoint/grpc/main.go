package main

import (
	gateway "github.com/noguchi-hiroshi/grpc-gateway-test/proto"
	"github.com/noguchi-hiroshi/grpc-gateway-test/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":4949"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	model := user.NewModel()
	srv := user.NewService(model)
	gateway.RegisterUserServiceServer(s, srv)
	log.Println("start grpc server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
