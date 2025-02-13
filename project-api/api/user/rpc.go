package user

import (
	"github.com/WildWeber-xiaocheng/ms_project/project-api/config"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/discovery"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/logs"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/user/login"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var LoginServiceClient login.LoginServiceClient

func InitRpcUserClient() {
	etcdRegister := discovery.NewResolver(config.Conf.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.Dial("etcd:///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = login.NewLoginServiceClient(conn)
}
