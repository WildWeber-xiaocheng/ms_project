package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-api/router"
)

// Route 接口的实现类
type RouterUser struct {
}

func init() {
	log.Println("init user router")
	ru := &RouterUser{}
	router.RegisterR(ru)
}

func (*RouterUser) Register(r *gin.Engine) {
	InitRpcUserClient() //初始化grpc客户端连接
	user := New()
	r.POST("/project/login/getCaptcha", user.GetCaptcha)
	r.POST("/project/login/register", user.Register)
}
