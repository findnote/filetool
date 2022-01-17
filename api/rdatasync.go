package main

import (
	"flag"
	"fmt"

	"mwp3000/api/config"
	"mwp3000/api/internal/handler"
	"mwp3000/api/internal/svc"
	"mwp3000/database"
	"mwp3000/pkg/redis/redisapi"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/rdata-sync-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// Setup Db(Sqlite)
	database.SetupDatabase()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	// 启动redis Listen节点
	go func() {
		redisapi.InitP3000RedisListen(c)
	}()

	server.Start()
}
