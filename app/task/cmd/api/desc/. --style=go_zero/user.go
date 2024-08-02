package main

import (
	"flag"
	"fmt"

	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/config"
	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/handler"
	"calendarAttempt3/app/task/cmd/api/desc/. --style=go_zero/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
