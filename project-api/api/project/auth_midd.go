package project

import (
	common "github.com/WildWeber-xiaocheng/ms_project/project-common"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/errs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// 不需要做授权操作的Url
var ignores = []string{
	"project/login/register",
	"project/login",
	"project/login/getCaptcha",
	"project/organization",
	"project/auth/apply",
}

func Auth() func(*gin.Context) {
	return func(c *gin.Context) {
		zap.L().Info("开始做授权认证")
		result := &common.Result{}
		//当用户登录认证通过，获取到用户信息，查询用户权限所拥有的节点信息
		//根据请求的uri路径 进行匹配 判断此URL是否在用户的授权列表中
		uri := c.Request.RequestURI

		//todo
		//还应该加一个判断，判断访问节点那里对接口是否做了权限控制，如果没有控制，则直接return

		for _, v := range ignores {
			if strings.Contains(uri, v) {
				c.Next()
				return
			}
		}

		a := NewAuth()
		nodes, err := a.GetAuthNodes(c)
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}

		for _, v := range nodes {
			if strings.Contains(uri, v) {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusOK, result.Fail(403, "无操作权限"))
		c.Abort()
	}
}
