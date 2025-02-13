package user

import (
	"github.com/WildWeber-xiaocheng/ms_project/project-api/api/midd"
	"github.com/WildWeber-xiaocheng/ms_project/project-api/api/rpc"
	"github.com/WildWeber-xiaocheng/ms_project/project-api/router"
	"github.com/gin-gonic/gin"
	"log"
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
