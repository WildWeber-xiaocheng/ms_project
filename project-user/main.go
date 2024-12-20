package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project-common"
	"test.com/project-user/config"
	"test.com/project-user/router"
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
	srv.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, stop)
}
