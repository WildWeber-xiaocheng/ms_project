package login_service_v1

import (
	"context"
	"go.uber.org/zap"
	"log"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-user/pkg/dao"
	"test.com/project-user/pkg/model"
	"test.com/project-user/pkg/repo"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	Cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		Cache: dao.Rc,
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	//1、获取参数
	//mobile := ctx.PostForm("mobile") 原来是通过gin来直接获取参数
	mobile := msg.Mobile
	//2、校验参数
	if ok := common.VerifyMobile(mobile); !ok {
		//ctx.JSON(http.StatusOK, rsp.Fail(model.LoginMobileNotLegal, "手机号不合法！"))
		//return nil, errors.New("手机号不合法")
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	//3、生成验证码 (也可以搞一个随机)
	code := "123456"
	//4、调用短信平台来发送验证码 这里放入go协程中执行，这样接口可以快速响应
	go func() {
		time.Sleep(2 * time.Second) //模拟短信发送延迟
		//log.Println("短信平台调用成功，发送短信")
		zap.L().Info("短信平台调用成功，发送短信")
		//5、存储验证码到redis中，过期时间为15min
		c, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelFunc()
		err := ls.Cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Println("验证码存入redis出错,cause by :", err)
		}
		log.Println("将手机号和验证码存入redis成功：REGISTER_%s:%s", mobile, code)
	}()
	//6、响应 正常真实情况不会返回code，所以上面的发送验证码会用go协程，因为不用返回数据，所以用协程处理发送验证码来快速响应本接口
	//ctx.JSON(http.StatusOK, rsp.Success(code))
	return &CaptchaResponse{Code: code}, nil
}
