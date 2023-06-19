package echoserver

import (
	"context"

	"google.golang.org/grpc"

	es "easy-gs/easyserver"
	pb "easy-gs/pb/github.com/akiuw/echo"
)

// server is used to implement helloworld.GreeterServer.
type EchoServer struct {
	pb.UnimplementedEchoServerServer
}

// SayHello implements helloworld.GreeterServer
func (es *EchoServer) Echo(ctx context.Context, in *pb.Echodata) (*pb.Echodata, error) {
	return &pb.Echodata{Data: in.Data}, nil
}

func (es *EchoServer) BeforeRun(_ *es.ServiceOpt) {
}

func (es *EchoServer) Run(opt *es.ServiceOpt, s grpc.ServiceRegistrar) {
	pb.RegisterEchoServerServer(s, &EchoServer{})
}
