package main

import (
	gs "easy-gs/easyserver"
	echo "easy-gs/echoserver"
	hello "easy-gs/helloserver"
)

func main() {
	es := &gs.EasyServer{}
	es = es.BuildServer(&hello.HelloServer{}).BuildServer(&echo.EchoServer{})
	es.Serve()
}
