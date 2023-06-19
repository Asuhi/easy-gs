package main

var GServerFlags *ServerFlags

type ServerFlags struct {
	Roles      []string
	ConfigPath string
}

func (sf *ServerFlags) GetPath() string {
	return sf.ConfigPath
}
func (sf *ServerFlags) GetConfigType() string {
	return "json"
}
func (sf *ServerFlags) GetRoles() []string {
	return sf.Roles
}
