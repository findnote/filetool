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

	Subscribe("CYGBase:RTDB-CYGBase:modify")
}

func TestPublish(t *testing.T) {
	flag.Parse()
	var c config.Config
	co.MustLoad(*path, &c)

	conf = c
	host := conf.Message.Host
	port := conf.Message.Port
	addr = host + ":" + port
	SessId = "5707483A-07F0-47D9-9A31-E63972CB9E7A"
	data := `{
    "name": "",
    "node": "",
    "topic": "",
    "sn": "",
    "data": "[{\"key\":\"CYGDW:Hash:Device:1:161\",\"tvModifys\":[{\"t\":\"yx\",\"sv\":\"0\",\"dv\":\"1\",\"time\":\"\"}]}]"
}`

	Publish([]byte(data), "CYGBase:RTDB-CYGBase:modify")
}
