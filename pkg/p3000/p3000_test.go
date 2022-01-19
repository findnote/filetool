package p3000_test

import (
	"flag"
	"mwp3000/api/config"
	"testing"

	"github.com/tal-tech/go-zero/core/conf"
)

var configFile = flag.String("f", "../../etc/mwp3000-api.yaml", "the config file")

func TestOrdersPostToServer(t *testing.T) {
	LoadConfig()

	// var data []interface{}
	// data = append(data, "testing")

	// p3000.OrdersPostToServer(data)
}

func LoadConfig() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// setting config
	// config.SetConfig(c)
}
