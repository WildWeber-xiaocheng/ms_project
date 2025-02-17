package main

import (
	_ "github.com/WildWeber-xiaocheng/ms_project/project-api/api"
	"github.com/WildWeber-xiaocheng/ms_project/project-api/api/midd"
	"github.com/WildWeber-xiaocheng/ms_project/project-api/config"
	_ "github.com/WildWeber-xiaocheng/ms_project/project-api/docs"
	"github.com/WildWeber-xiaocheng/ms_project/project-api/router"
	srv "github.com/WildWeber-xiaocheng/ms_project/project-common"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title project-api
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @host 这里写接口服务的host
// @BasePath 这里写base path
func main() {
	r := gin.Default()
	r.Use(midd.RequestLog())
	r.StaticFS("/upload", http.Dir("upload"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.InitRouter(r)
	//开启pprof 默认的访问路径是/debug/pprof
	pprof.Register(r)
	srv.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, nil)
}
