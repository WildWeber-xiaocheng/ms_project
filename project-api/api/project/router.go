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
	h := New()
	group := r.Group("/project")
	group.Use(midd.TokenVerify())
	group.POST("/index", h.index)
	group.POST("/project/selfList", h.myProjectList)
	group.POST("/project", h.myProjectList) //和'/selfList'共用一个处理函数
	group.POST("/project_template", h.projectTemplate)
	group.POST("/project/save", h.projectSave)
	group.POST("/project/read", h.readProject)
}
