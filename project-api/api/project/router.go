package project

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-api/api/midd"
	"test.com/project-api/router"
)

// Route 接口的实现类
type RouterProject struct {
}

func init() {
	log.Println("init project router")
	ru := &RouterProject{}
	router.RegisterR(ru)
}

func (*RouterProject) Register(r *gin.Engine) {
	InitRpcProjectClient() //初始化grpc客户端连接
	user := New()
	group := r.Group("/project/index")
	group.Use(midd.TokenVerify())
	group.POST("", user.index)
	group1 := r.Group("/project/project")
	group1.Use(midd.TokenVerify())
	group1.POST("selfList", user.myProjectList)
}
