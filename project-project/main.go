package main

import (
	srv "github.com/WildWeber-xiaocheng/ms_project/project-common"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/config"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	//初始化rpc调用
	router.InitUserRpc()
	//grpc服务注册
	gc := router.RegisterGrpc()
	//grpc服务注册到etcd中
	router.RegisterEtcdServer()
	stop := func() {
		gc.Stop()
	}

	srv.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, stop)
}
