package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project-common"
	_ "test.com/project-user/api"
	"test.com/project-user/config"
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	//grpc服务注册
	gc := router.RegisterGrpc()
	stop := func() {
		gc.Stop()
	}
	srv.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, stop)
}
