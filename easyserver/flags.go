package easyserver

import (
	"flag"
	"strings"
)

var GServerFlags *ServerFlags

type ServerFlags struct {
	Roles      []string
	ConfigPath string
}

func init() {
	GServerFlags = new(ServerFlags)

	var roles string
	flag.StringVar(&roles, "roles", "all", "started with service role, all: start as all role in config")
	flag.StringVar(&GServerFlags.ConfigPath, "config", "./config/config.json", "config path")
	flag.Parse()
	GServerFlags.Roles = strings.Split(roles, ";")
}
