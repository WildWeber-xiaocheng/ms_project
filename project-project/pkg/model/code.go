package model

import (
	"test.com/project-common/errs"
)

var (
	RedisError  = errs.NewError(999, "redis错误")
	DBError     = errs.NewError(998, "db错误")
	ParamsError = errs.NewError(401, "参数错误")
	//1010表示业务逻辑 10表示user 10表示login
	NoLegalMobile         = errs.NewError(10102001, "手机号不合法")
	CaptchaNotExist       = errs.NewError(10102002, "验证码不存在/已过期")
	CaptchaError          = errs.NewError(10102003, "验证码错误")
	EmailExist            = errs.NewError(10102004, "邮箱已经存在")
	AccountExist          = errs.NewError(10102005, "账号已经存在")
	MobileExist           = errs.NewError(10102006, "手机号已经存在")
	AccountAndPwdError    = errs.NewError(10102007, "账号或密码不正确")
	TaskNameNotNull       = errs.NewError(20102001, "任务标题不能为空")
	TaskStagesNotNull     = errs.NewError(20102002, "任务步骤不存在")
	ProjectAlreadyDeleted = errs.NewError(20102003, "项目已经删除")
)
