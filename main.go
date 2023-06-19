package main

import (
	gs "easy-gs/easyserver"
	echo "easy-gs/echoserver"
	hello "easy-gs/helloserver"
	"flag"
	"strings"
)

func main() {
	GServerFlags := &ServerFlags{}

	var roles string
	flag.StringVar(&roles, "roles", "all", "started with service role, all: start as all role in config")
	flag.StringVar(&GServerFlags.ConfigPath, "config", "./config/config.json", "config path")
	flag.Parse()
	GServerFlags.Roles = strings.Split(roles, ";")

	es := &gs.EasyServer{}

	es.SetFlags(GServerFlags)
	es.LoadFlagConfig()

	es = es.BuildServer("hello", &hello.HelloServer{}).BuildServer("echo", &echo.EchoServer{})
	es.Serve()
}
