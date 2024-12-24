package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test.com/project-api/pkg/model/user"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/user/login"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (h *HandlerUser) GetCaptcha(ctx *gin.Context) {
	rsp := &common.Result{}
	//1、获取参数
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	captcha, err := LoginServiceClient.GetCaptcha(c, &login.CaptchaMessage{Mobile: mobile})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, rsp.Fail(code, msg))
	}
	ctx.JSON(http.StatusOK, rsp.Success(captcha.Code))
}

func (h *HandlerUser) Register(c *gin.Context) {
	result := &common.Result{}
	//1. 接收参数
	var req user.RegisterReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	//2. 参数校验
	if err := req.Verify(); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//处理业务
	//方法一
	//msg := &login.RegisterMessage{
	//	Name:     req.Name,
	//	Email:    req.Email,
	//	Mobile:   req.Mobile,
	//	Password: req.Password,
	//	Captcha:  req.Captcha,
	//}

	//方法二 使用copier
	msg := &login.RegisterMessage{}
	err = copier.Copy(msg, req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy有误"))
		return
	}
	//3. 调用user grpc服务，获取相应
	_, err = LoginServiceClient.Register(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//4. 返回结果
	c.JSON(http.StatusOK, result.Success(""))
}

func (h *HandlerUser) Login(c *gin.Context) {
	result := &common.Result{}
	//1. 接收参数
	var req user.LoginReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	//2. 调用user grpc 完成登录
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	msg := &login.LoginMessage{}
	err = copier.Copy(msg, req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy有误"))
		return
	}
	loginRsp, err := LoginServiceClient.Login(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	rsp := &user.LoginRsp{}
	err = copier.Copy(rsp, loginRsp)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy有误"))
		return
	}
	//4. 返回结果
	c.JSON(http.StatusOK, result.Success(rsp))
}
