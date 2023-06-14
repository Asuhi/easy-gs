package easyserver

// easyserver: framework
// it has all of other servers needed things eg: tools middleware etc...

// gameserver: players frame sync play etc..
// roomserver: players matchmaking group stats etc...
// relationserver: frend etc...
// chatserver: chat
// loginserver: login get token...
// gateway: http -> grpc
// notifyserver: tcp ---- client

type EasyServer struct {
	ServiceName string
	// tracer
	// rpc service instance
	// http server instance
	// messagequeue
	// etcd client
}
