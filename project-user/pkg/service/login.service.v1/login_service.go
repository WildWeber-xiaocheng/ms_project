package login_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	common "test.com/project-common"
	"test.com/project-common/encrypts"
	"test.com/project-common/errs"
	"test.com/project-grpc/user/login"
	"test.com/project-user/internal/dao"
	"test.com/project-user/internal/data/member"
	"test.com/project-user/internal/data/organization"
	"test.com/project-user/internal/database"
	"test.com/project-user/internal/database/tran"
	"test.com/project-user/internal/repo"
	"test.com/project-user/pkg/model"
	"time"
)

type LoginService struct {
	login.UnimplementedLoginServiceServer
	cache            repo.Cache
	memberRepo       repo.MemberRepo
	organizationRepo repo.OrganizationRepo
	transaction      tran.Transaction
}

func New() *LoginService {
	return &LoginService{
		cache:            dao.Rc,
		memberRepo:       dao.NewMemberDao(),
		organizationRepo: dao.NewOrganizationDao(),
		transaction:      dao.NewTransaction(),
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *login.CaptchaMessage) (*login.CaptchaResponse, error) {
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
		err := ls.cache.Put(c, model.RegisterRedisKey+mobile, code, 15*time.Minute)
		if err != nil {
			log.Println("验证码存入redis出错,cause by :", err)
		}
		log.Println("将手机号和验证码存入redis成功：REGISTER_%s:%s", mobile, code)
	}()
	//6、响应 正常真实情况不会返回code，所以上面的发送验证码会用go协程，因为不用返回数据，所以用协程处理发送验证码来快速响应本接口
	//ctx.JSON(http.StatusOK, rsp.Success(code))
	return &login.CaptchaResponse{Code: code}, nil
}

func (ls *LoginService) Register(ctx context.Context, msg *login.RegisterMessage) (*login.RegisterResponse, error) {
	//1. 可以再次检验参数（这里就省略了）
	//2. 校验验证码
	c := context.Background()
	redisCode, err := ls.cache.Get(c, model.RegisterRedisKey+msg.Mobile)

	//redis中查不到key时也返回错误
	if err == redis.Nil {
		return nil, errs.GrpcError(model.CaptchaNotExist)
	}
	if err != nil {
		zap.L().Error("Register redis get error", zap.Error(err))
		return nil, errs.GrpcError(model.RedisError)
	}
	if redisCode != msg.Captcha {
		return nil, errs.GrpcError(model.CaptchaError)
	}
	//3. 校验业务逻辑（例如：邮箱是否被注册 账号是否被注册 手机号是否被注册）
	//邮箱
	exist, err := ls.memberRepo.GetMemberByEmail(c, msg.Email)
	if err != nil {
		zap.L().Error("Register db get error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.EmailExist)
	}
	//账号
	exist, err = ls.memberRepo.GetMemberByEmail(c, msg.Name)
	if err != nil {
		zap.L().Error("Register db get error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.AccountExist)
	}
	//手机号
	exist, err = ls.memberRepo.GetMemberByEmail(c, msg.Mobile)
	if err != nil {
		zap.L().Error("Register db get error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.MobileExist)
	}
	//4. 执行业务 将数据存入member表，生成一个数据，存入organization组织表
	pwd := encrypts.Md5(msg.Password)
	mem := &member.Member{
		Account:       msg.Name,
		Password:      pwd,
		Name:          msg.Name,
		Mobile:        msg.Mobile,
		Email:         msg.Email,
		CreateTime:    time.Now().UnixMilli(),
		LastLoginTime: time.Now().UnixMilli(),
		Status:        model.Normal,
	}
	err = ls.transaction.Action(func(conn database.DbConn) error {
		err = ls.memberRepo.SaveMember(conn, c, mem)
		if err != nil {
			zap.L().Error("Register db SaveMember error", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		org := &organization.Organization{
			Name:       mem.Name + "个人组织",
			MemberId:   mem.Id,
			CreateTime: time.Now().UnixMilli(),
			Personal:   model.Personal,
			Avatar:     "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5",
		}
		err = ls.organizationRepo.SaveOrganization(conn, c, org)
		if err != nil {
			zap.L().Error("register SaveOrganization db err", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		return nil
	})
	//5. 返回结果
	return &login.RegisterResponse{}, err
}

func (ls *LoginService) Login(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	c := context.Background()
	//1. 去数据库查询账号和密码，看是否正确
	pwd := encrypts.Md5(msg.Password) //因为在注册的时候进行md5加密了，所以这里也需要加密一下
	mem, err := ls.memberRepo.FindMember(c, msg.Account, pwd)
	if err != nil {
		//if errors.Is(err, model.DataIsNull) {
		//	return &LoginResponse{}, model.AccountAndPwdError
		//}
		zap.L().Error("Login FindMember db fail", zap.Error(err))
		return &login.LoginResponse{}, errs.GrpcError(model.DBError)
	}
	if mem == nil {
		return &login.LoginResponse{}, errs.GrpcError(model.AccountAndPwdError)
	}

	memMsg := &login.MemberMessage{}
	err = copier.Copy(memMsg, mem)
	//2. 登录成功 根据用户id查询组织
	orgs, err := ls.organizationRepo.FindOrganizationByMemId(c, mem.Id)
	if err != nil {
		zap.L().Error("Login FindOrganizationByMemId db fail", zap.Error(err))
		return &login.LoginResponse{}, errs.GrpcError(model.DBError)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)

	//jwtToken := jwts.CreateToken(strconv.FormatInt(member.Id, 10), "msproject", 3600*24*7*time.Second, 3600*24*14*time.Second)
	//tokenList := &TokenMessage{
	//	AccessToken:    jwtToken.AccessToken,
	//	RefreshToken:   jwtToken.RefreshToken,
	//	AccessTokenExp: jwtToken.AccessExp.Milliseconds() / 1000,
	//	TokenType:      "bearer",
	//}

	return &login.LoginResponse{
		Member:           memMsg,
		OrganizationList: orgsMessage,
	}, nil
}
