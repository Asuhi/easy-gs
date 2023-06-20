package echoserver

import (
	"context"

	"google.golang.org/grpc"

	echo "akiuw.golang.com/easy-gs/pb/echo"
	esv "akiuw.golang.com/easyserver"
)

// server is used to implement helloworld.GreeterServer.
type EchoServer struct {
	echo.UnimplementedEchoServerServer
}

// SayHello implements helloworld.GreeterServer
func (es *EchoServer) Echo(ctx context.Context, in *echo.Echodata) (*echo.Echodata, error) {
	return &echo.Echodata{Data: in.Data}, nil
}

func (es *EchoServer) BeforeRun(_ *esv.ServiceOpt) {
}

func (es *EchoServer) Run(opt *esv.ServiceOpt, s grpc.ServiceRegistrar) {
	echo.RegisterEchoServerServer(s, &EchoServer{})
}
