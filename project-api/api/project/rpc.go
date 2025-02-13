package project

import (
	"github.com/WildWeber-xiaocheng/ms_project/project-api/config"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/discovery"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/logs"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/account"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/auth"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/department"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/menu"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/project"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var ProjectServiceClient project.ProjectServiceClient
var TaskServiceClient task.TaskServiceClient
var AccountServiceClient account.AccountServiceClient
var DepartmentServiceClient department.DepartmentServiceClient
var AuthServiceClient auth.AuthServiceClient
var MenuServiceClient menu.MenuServiceClient

func InitRpcProjectClient() {
	etcdRegister := discovery.NewResolver(config.Conf.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.Dial("etcd:///project", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ProjectServiceClient = project.NewProjectServiceClient(conn)
	TaskServiceClient = task.NewTaskServiceClient(conn)
	AccountServiceClient = account.NewAccountServiceClient(conn)
	DepartmentServiceClient = department.NewDepartmentServiceClient(conn)
	AuthServiceClient = auth.NewAuthServiceClient(conn)
	MenuServiceClient = menu.NewMenuServiceClient(conn)
}
