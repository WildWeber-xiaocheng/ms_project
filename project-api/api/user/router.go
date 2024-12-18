package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-api/router"
)

func init() {
	log.Println("init user router")
	router.RegisterR(&RouterUser{})
}

// Route 接口的实现类
type RouterUser struct {
}

func (*RouterUser) Register(r *gin.Engine) {
	InitRpcUserClient() //初始化grpc客户端连接
	user := New()
	r.POST("/project/login/getCaptcha", user.GetCaptcha)
}
