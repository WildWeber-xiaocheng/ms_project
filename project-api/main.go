package main

import (
	"github.com/gin-gonic/gin"
	_ "test.com/project-api/api"
	"test.com/project-api/api/midd"
	"test.com/project-api/config"
	"test.com/project-api/router"
	srv "test.com/project-common"
)

func main() {
	r := gin.Default()
	r.Use(midd.RequestLog())
	router.InitRouter(r)
	srv.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, nil)
}
