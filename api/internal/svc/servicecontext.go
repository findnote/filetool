package svc

import (
	"mwp3000/api/config"
	"mwp3000/pkg/p3000"
)

type ServiceContext struct {
	Config    config.Config
	P3000Sync *p3000.P3000Conn
}

func NewServiceContext(c config.Config) *ServiceContext {
	p3000Conn := p3000.NewP3000Client(c)
	return &ServiceContext{
		Config:    c,
		P3000Sync: p3000Conn,
	}
}
