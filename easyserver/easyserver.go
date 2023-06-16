package easyserver

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"path"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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
	BeforeRun(*ServiceOpt)
	Run(context.Context, *ServiceOpt, grpc.ServiceRegistrar)
	// Config()
	// Tracer()
}

// a node has some roles use -roles="" to specify
// eg: ./easy-gos -roles="hello;chat"

// use -config="" to specify config path
// config file is a json file

type Server struct {
	Opt    *ServiceOpt
	Server IServer
	DB     *sqlx.DB
	Viper  *viper.Viper
}

type EasyServer struct {
	Servers map[string]*Server // k:servicename
	// rpc service instance
	// http server instance
	// messagequeue
}

func (es *EasyServer) BuildServer(server IServer) *EasyServer {

	viper.SetConfigName(path.Base(GServerFlags.ConfigPath)) // name of config file (without extension)
	viper.SetConfigType("json")                             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path.Dir(GServerFlags.ConfigPath))  // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	f, err := os.Open(GServerFlags.ConfigPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	opts := Options{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&opts); err != nil {
		log.Fatalln(err)
	}

	if len(es.Servers) == 0 {
		es.Servers = make(map[string]*Server)
	}
	for _, v := range opts.Services {
		var db *sqlx.DB
		if v.Database != "" {
			db, err := sqlx.Open("mysql", v.Database)
			if err != nil {
				log.Fatalln(err)
			}
			db.SetMaxOpenConns(v.PoolConns)
			db.SetMaxIdleConns(v.PoolConns)
		}

		es.Servers[v.Name] = &Server{
			Opt:    v,
			Server: server,
			Viper:  viper.GetViper(),
			DB:     db,
		}
	}

	return es
}

func (es *EasyServer) GrpcServe(opt *ServiceOpt) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", opt.ListenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return grpc.NewServer(), lis
}

func (es *EasyServer) Serve() {
	ctx := context.Background()
	for _, role := range GServerFlags.Roles {
		if s, ok := es.Servers[role]; ok {
			go func(ctx context.Context, s *Server) {

				cancelCtx, f := context.WithCancel(ctx)
				s.Server.BeforeRun(s.Opt)
				g, lis := es.GrpcServe(s.Opt)
				s.Server.Run(cancelCtx, s.Opt, g)
				log.Printf("%s service listening at %v", s.Opt.Name, lis.Addr())
				if err := g.Serve(lis); err != nil {
					log.Fatalf("failed to serve: %v", err)
				}
				f()
				cancelCtx.Done()
			}(ctx, s)
		}
	}
	<-ctx.Done()
}
