package project

import (
	"github.com/gin-gonic/gin"
	"log"
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
	InitRpcUserClient() //初始化grpc客户端连接
	user := New()
	r.POST("/project/index", user.index)
}
