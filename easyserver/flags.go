package easyserver

type IServerFlags interface {
	GetPath() string
	GetConfigType() string
	GetRoles() []string
}

func (es *EasyServer) SetFlags(f IServerFlags) {
	es.Flags = f
}
