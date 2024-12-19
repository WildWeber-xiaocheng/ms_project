package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-user/pkg/dao"
	"test.com/project-user/pkg/repo"
	loginServiceV1 "test.com/project-user/pkg/service/login.service.v1"
	"time"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc,
	}
}

func (h *HandlerUser) GetCaptcha(ctx *gin.Context) {
	rsp := &common.Result{}
	//1、获取参数
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	captcha, err := LoginServiceClient.GetCaptcha(c, &loginServiceV1.CaptchaMessage{Mobile: mobile})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, rsp.Fail(code, msg))
	}
	ctx.JSON(http.StatusOK, rsp.Success(captcha.Code))
}
