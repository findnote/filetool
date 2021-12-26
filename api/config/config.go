package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf

	// P3000
	P3000 struct {
		Host string
		Port int
	}

	// Gis服务器地址（ws）
	// ws://ip:port/securityWs/thirdParty/{token}
	GisPortal struct {
		Token string
		Host  string
		Port  int
	}

	// p3000redis
	P3000Redis struct {
		Host string
		Port int
		Pwd  string
		DB   int
	}

	// gisdemo
	GisDemo struct {
		Enable   bool
		Interval int
	}
}
