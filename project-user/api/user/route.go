package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-user/router"
)

func init() {
	log.Println("init router user")
	router.RegisterR(&RouterUser{})
}

// Route 接口的实现类
type RouterUser struct {
}

func (*RouterUser) Register(r *gin.Engine) {
	user := New()
	r.POST("/project/login/getCaptcha", user.GetCaptcha)
}
