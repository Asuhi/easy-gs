package easyserver

type Options struct {
	Services []*ServiceOpt `json:"services"`
	MQOption *NSQOpt       `json:"nsq"`
}

type NSQOpt struct {
	Address    string `json:"address"`
	LookupAddr string `json:"lookupaddr"`
}

type ServiceOpt struct {
	Name       string `json:"name"`
	ListenPort string `json:"listen_port"`
	Redis      string `json:"redis"`
	Database   string `json:"database"`
	PoolConns  int    `json:"pool_conns"`
}
