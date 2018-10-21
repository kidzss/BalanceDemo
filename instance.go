package balance

import (
	"strconv"
)

type Instance struct {
	host string
	port int
}

func NewInstance(host string,port int) *Instance {
	return &Instance{
		host:host,
		port:port,
	}
}

//定义Instance结构体的方法
func (p *Instance)GetPort() int {
	return p.port
}

func (p *Instance)GetHost() string  {
	return p.host
}

func (p *Instance) String() string {
	return p.host + ":" +strconv.Itoa(p.port)
}