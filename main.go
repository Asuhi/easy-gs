package main

import (
	gs "easy-gs/easyserver"
	hello "easy-gs/helloserver"
)

func main() {

	es := &gs.EasyServer{}
	es = es.BuildServer(&hello.HelloServer{})
	es.Serve()
}
