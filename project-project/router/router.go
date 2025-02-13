package router

import (
	"github.com/WildWeber-xiaocheng/ms_project/project-common/discovery"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/logs"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/account"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/auth"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/department"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/menu"
	project "github.com/WildWeber-xiaocheng/ms_project/project-grpc/project"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/task"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/config"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/interceptor"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/rpc"
	account_service_v1 "github.com/WildWeber-xiaocheng/ms_project/project-project/pkg/service/account.service.v1"
	auth_service_v1 "github.com/WildWeber-xiaocheng/ms_project/project-project/pkg/service/auth.service.v1"
	department_service_v1 "github.com/WildWeber-xiaocheng/ms_project/project-project/pkg/service/department.service.v1"
	menu_service_v1 "github.com/WildWeber-xiaocheng/ms_project/project-project/pkg/service/menu.service.v1"
	project_service_v1 "github.com/WildWeber-xiaocheng/ms_project/project-project/pkg/service/project.service.v1"
	task_service_v1 "github.com/WildWeber-xiaocheng/ms_project/project-project/pkg/service/task.service.v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
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

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.Conf.GC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			project.RegisterProjectServiceServer(g, project_service_v1.New())
			task.RegisterTaskServiceServer(g, task_service_v1.New())
			account.RegisterAccountServiceServer(g, account_service_v1.New())
			department.RegisterDepartmentServiceServer(g, department_service_v1.New())
			auth.RegisterAuthServiceServer(g, auth_service_v1.New())
			menu.RegisterMenuServiceServer(g, menu_service_v1.New())
		}}
	s := grpc.NewServer(interceptor.New().Cache())
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}

func RegisterEtcdServer() {
	etcdRegister := discovery.NewResolver(config.Conf.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	info := discovery.Server{
		Name:    config.Conf.GC.Name,
		Addr:    config.Conf.GC.Addr,
		Version: config.Conf.GC.Version,
		Weight:  config.Conf.GC.Weight,
	}
	r := discovery.NewRegister(config.Conf.EtcdConfig.Addrs, logs.LG)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}

func InitUserRpc() {
	rpc.InitRpcUserClient()
}
