package main

import (
	gs "easy-gs/easyserver"
	hello "easy-gs/helloserver"
)

func main() {

	es := &gs.EasyServer{}

	opt := &gs.ServerOpt{
		ServiceName: "hello",
	}
	es = es.BuildServer(opt, hello.HelloServer)
	es.Serve()
}
