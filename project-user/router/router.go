package router

import (
	"github.com/WildWeber-xiaocheng/ms_project/project-common/discovery"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/logs"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/user/login"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/config"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/interceptor"
	loginServiceV1 "github.com/WildWeber-xiaocheng/ms_project/project-user/pkg/service/login.service.v1"
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
			login.RegisterLoginServiceServer(g, loginServiceV1.New())
		}}
	//in := grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//	if info.FullMethod != "/login.service.v1.LoginService/MyOrgList" {
	//		return handler(ctx, req)
	//	}
	//	fmt.Println("请求之前")
	//	rsp, err := handler(ctx, req)
	//	fmt.Println("请求之后")
	//	return rsp, err
	//})
	cacheInterceptor := interceptor.New()
	s := grpc.NewServer(cacheInterceptor.Cache())
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
