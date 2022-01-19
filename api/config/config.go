package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf

	// P3000
	P3000 struct {
		Host string
		Port int
	}

	// p3000redis
	P3000Redis struct {
		Host string
		Port int
		Pwd  string
		DB   int
	}
}
