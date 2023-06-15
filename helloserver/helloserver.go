package helloserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "easy-gs/pb/github.com/akiuw/hello"
)

// server is used to implement helloworld.GreeterServer.
type HelloServer struct {
	pb.UnimplementedHelloServerServer
}

// SayHello implements helloworld.GreeterServer
func (hs *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (hs *HelloServer) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServerServer(s, &HelloServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
