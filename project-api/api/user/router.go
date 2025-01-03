package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-api/api/midd"
	"test.com/project-api/api/rpc"
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
	rpc.InitRpcUserClient() //初始化grpc客户端连接
	user := New()
	r.POST("/project/login/getCaptcha", user.GetCaptcha)
	r.POST("/project/login/register", user.Register)
	r.POST("/project/login", user.Login)
	org := r.Group("/project/organization")
	org.Use(midd.TokenVerify())
	org.POST("/_getOrgList", user.myOrgList)
}
