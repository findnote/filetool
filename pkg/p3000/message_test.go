package p3000

import (
	"flag"
	co "github.com/tal-tech/go-zero/core/conf"
	"mwp3000/api/config"
	"testing"
)

var path = flag.String("path", "../../etc/rdata-sync-api.yaml", "config")

func TestSubscribe(t *testing.T) {
	flag.Parse()
	var c config.Config
	co.MustLoad(*path, &c)

	conf = c
	host := conf.Message.Host
	port := conf.Message.Port
	addr = host + ":" + port
	SessId = "E7666EBA-1F47-4784-A598-FFA61FDDCFBF"

	Subscribe()
}

func TestPublish(t *testing.T) {
	flag.Parse()
	var c config.Config
	co.MustLoad(*path, &c)

	conf = c
	host := conf.Message.Host
	port := conf.Message.Port
	addr = host + ":" + port
	SessId = "1718989F-EE49-40E9-9DB2-FCC6C7C75286"

	data := `{
    "2333": "233323"
}`
	Publish([]byte(data))
}
