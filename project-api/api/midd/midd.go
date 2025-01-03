package midd

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"test.com/project-api/api/rpc"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/user/login"
	"time"
)

func TokenVerify() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		result := &common.Result{}
		//1. 从header中获取token
		token := c.GetHeader("Authorization")
		//2. 调用user服务进行token认证
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
		defer cancelFunc()
		response, err := rpc.LoginServiceClient.TokenVerify(ctx, &login.LoginMessage{Token: token})
		//3. 处理结果，认证通过，将信息放入gin的上下文，失败则返回未登录
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}
		c.Set("memberId", response.Member.Id)
		c.Set("memberName", response.Member.Name)
		c.Next()
	}
}
