package main

import (
	srv "github.com/WildWeber-xiaocheng/ms_project/project-common"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/config"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	//grpc服务注册
	gc := router.RegisterGrpc()
	//grpc服务注册到etcd中
	router.RegisterEtcdServer()
	stop := func() {
		gc.Stop()
	}
	//gin优雅启停
	srv.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, stop)
}
