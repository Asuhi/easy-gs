package main

import (
	gs "akiuw.golang.com/easyserver"

	echo "akiuw.golang.com/easy-gs/echoserver"
	hello "akiuw.golang.com/easy-gs/helloserver"

	"flag"
	"strings"
)

func main() {
	gServerFlags := &ServerFlags{}

	var roles string
	flag.StringVar(&roles, "roles", "all", "started with service role, all: start as all role in config")
	flag.StringVar(&gServerFlags.ConfigPath, "config", "./config/config.json", "config path")
	flag.Parse()
	gServerFlags.Roles = strings.Split(roles, ";")

	es := &gs.EasyServer{}

	es.LoadFlagConfig(gServerFlags)

	es = es.BuildServer("hello", &hello.HelloServer{}).BuildServer("echo", &echo.EchoServer{})
	es.Serve()
}
