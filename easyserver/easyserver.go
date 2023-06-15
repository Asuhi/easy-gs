package easyserver

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// easyserver: framework
// it has all of other servers needed things eg: tools middleware etc...

// gameserver: players frame sync play etc..
// roomserver: players matchmaking group stats etc...
// relationserver: frend etc...
// chatserver: chat
// loginserver: login get token...
// gateway: http -> grpc
// notifyserver: tcp ---- client
type IServer interface {
	Run()
	// Config()
	// Tracer()
}

type ServerOpt struct {
	ServiceName string
	RedisAddr   string
	DBTarget    string // using tidb
}

type Server struct {
	Opt    *ServerOpt
	Server IServer
	DB     *sqlx.DB
}

type EasyServer struct {
	Servers []*Server
	// tracer
	// rpc service instance
	// http server instance
	// messagequeue
}

func (es *EasyServer) BuildServer(opt *ServerOpt, server IServer) *EasyServer {
	s := &Server{
		Opt:    opt,
		Server: server,
	}
	es.Servers = append(es.Servers, s)
	return es
}

func NewOption(serviceName, redisAddr, dbTarget string) *ServerOpt {
	return &ServerOpt{
		ServiceName: serviceName,
		RedisAddr:   redisAddr,
		DBTarget:    dbTarget,
	}
}

func (es *EasyServer) Serve() {
	ctx := context.Background()
	for _, s := range es.Servers {
		go func(ctx context.Context, s *Server) {
			s.Server.Run()
		}(ctx, s)
	}
}
