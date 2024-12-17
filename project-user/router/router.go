package router

import (
	"github.com/gin-gonic/gin"
)

// Router接口
type Router interface {
	Register(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Route(router Router, r *gin.Engine) {
	router.Register(r)
}

var routers []Router

func InitRouter(r *gin.Engine) {
	//router := New()
	//router.Route(&user.RouterUser{}, r)
	for _, ro := range routers {
		ro.Register(r)
	}
}

func RegisterR(ro ...Router) {
	routers = append(routers, ro...)
}
