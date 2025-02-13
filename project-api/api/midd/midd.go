package midd

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-api/api/rpc"
	common "github.com/WildWeber-xiaocheng/ms_project/project-common"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/errs"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/user/login"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 获取ip函数
func GetIp(c *gin.Context) string {
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}

func TokenVerify() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		result := &common.Result{}
		//1. 从header中获取token
		token := c.GetHeader("Authorization")
		//2. 调用user服务进行token认证
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
		defer cancelFunc()
		ip := GetIp(c)
		//todo
		//可以先去查询node表，确认不使用登录控制的接口，就不做登录认证
		response, err := rpc.LoginServiceClient.TokenVerify(ctx, &login.LoginMessage{Token: token, Ip: ip})
		//3. 处理结果，认证通过，将信息放入gin的上下文，失败则返回未登录
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}
		c.Set("memberId", response.Member.Id)
		c.Set("memberName", response.Member.Name)
		c.Set("organizationCode", response.Member.OrganizationCode)
		c.Next()
	}
}
