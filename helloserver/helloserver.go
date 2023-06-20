package helloserver

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	echo "akiuw.golang.com/easy-gs/pb/echo"
	hello "akiuw.golang.com/easy-gs/pb/hello"

	esv "akiuw.golang.com/easyserver"
)

// server is used to implement helloworld.GreeterServer.
type HelloServer struct {
	hello.UnimplementedHelloServerServer
}

// SayHello implements helloworld.GreeterServer
func (hs *HelloServer) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	conn, err := grpc.Dial("127.0.0.1:10003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	cli := echo.NewEchoServerClient(conn)
	out, err := cli.Echo(ctx, &echo.Echodata{
		Data: in.GetName(),
	})
	if err != nil {
		return nil, err
	}
	return &hello.HelloReply{Message: "Hello " + out.GetData()}, nil
}

func (hs *HelloServer) BeforeRun(_ *esv.ServiceOpt) {

}

func (hs *HelloServer) Run(opt *esv.ServiceOpt, s grpc.ServiceRegistrar) {
	hello.RegisterHelloServerServer(s, &HelloServer{})
}
