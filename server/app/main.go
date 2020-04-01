package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"

	"disconf/server/api"
	"disconf/server/model"
	"disconf/server/tcp"
	"disconf/server/util"
)

const (
	defaultConfigPath = "/Users/jiajianyun/go/src/disconf/server/etc/config.yaml"
)

func init() {
	var err error
	if err = util.InitConfig(defaultConfigPath); err != nil {
		goto exit
	}
	if err = util.InitDB(); err != nil {
		goto exit
	}
	if err = util.InitLogger(); err != nil {
		goto exit
	}
	return
exit:
	panic(err)
}

func main() {
	//启动tcp服务
	tcpaddr := fmt.Sprintf("%s:%s", util.G_conf.ServerConfig.IP, util.G_conf.ServerConfig.Port)
	server := tcp.NewServer(tcpaddr)
	if err := server.Run(); err != nil {
		panic(err)
	}
	//启动http服务
	e := echo.New()
	initRouter(e, server)
	go func() {
		defer func() {
			if err := recover(); err != nil {
			}
		}()
		addr := fmt.Sprintf("%s:%s", util.G_conf.ServerConfig.IP, util.G_conf.ServerConfig.HttpPort)
		if err := e.Start(addr); err != nil {
		}
	}()
	waitStop()
}

func initRouter(e *echo.Echo, s model.Server) {
	g := e.Group("/api/v1")
	app := &api.App{}
	g.POST("/createapp", app.CreateApp)
	cluster := &api.Cluster{}
	g.POST("/createcluster", cluster.CreateCluster)
	release := &api.Release{
		Server: s,
	}
	g.POST("/pubconfig", release.PubcConfig)
}

func waitStop() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-signalChan
}
