package helloserver

import (
	"context"
	"log"

	"google.golang.org/grpc"

	es "easy-gs/easyserver"
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

func (hs *HelloServer) BeforeRun(_ *es.ServiceOpt) {

}

func (hs *HelloServer) Run(ctx context.Context, opt *es.ServiceOpt, s grpc.ServiceRegistrar) {
	pb.RegisterHelloServerServer(s, &HelloServer{})
}
