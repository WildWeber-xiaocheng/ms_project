package login_service_v1

import (
	"context"
	"encoding/json"
	common "github.com/WildWeber-xiaocheng/ms_project/project-common"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/encrypts"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/errs"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/jwts"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/tms"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/user/login"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/config"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/dao"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/data/member"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/data/organization"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/database"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/database/tran"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/repo"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/pkg/model"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	"strconv"
	"strings"
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
	//todo
	//生成账户完毕后要添加一个账户，包含账户的授权角色：成员，新生成一个角色（成员）（如果没有），同时将此角色的Node生成
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
	memMsg.Code, _ = encrypts.EncryptInt64(mem.Id, model.AESKey)
	memMsg.LastLoginTime = tms.FormatByMill(mem.LastLoginTime)
	memMsg.CreateTime = tms.FormatByMill(mem.CreateTime)
	//2. 登录成功 根据用户id查询组织
	orgs, err := ls.organizationRepo.FindOrganizationByMemId(c, mem.Id)
	if err != nil {
		zap.L().Error("Login FindOrganizationByMemId db fail", zap.Error(err))
		return &login.LoginResponse{}, errs.GrpcError(model.DBError)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)
	for _, org := range orgsMessage {
		org.Code, _ = encrypts.EncryptInt64(org.Id, model.AESKey)
		org.OwnerCode = memMsg.Code
		org.CreateTime = tms.FormatByMill(organization.ToMap(orgs)[org.Id].CreateTime)
	}
	if len(orgs) > 0 { //对第一个组织进行加密
		memMsg.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
	}
	//3. 用jwt生成token
	memIdStr := strconv.FormatInt(mem.Id, 10)
	//尽管time.Second是int64，但是这里不能直接相乘
	exp := time.Duration(config.Conf.JwtConfig.AccessExp*3600*24) * time.Second
	rExp := time.Duration(config.Conf.JwtConfig.RefreshExp*3600*24) * time.Second
	token := jwts.CreateToken(memIdStr, exp, config.Conf.JwtConfig.AccessSecret,
		rExp, config.Conf.JwtConfig.RefreshSecret, msg.Ip)
	//可以给token做加密处理 增加安全性
	tokenList := &login.TokenMessage{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		AccessTokenExp: token.AccessExp,
		TokenType:      "bearer", //先固定为这个字段
	}
	//放入缓存
	go func() {
		memJson, _ := json.Marshal(mem)
		ls.cache.Put(c, model.Member+"::"+memIdStr, string(memJson), exp)
		orgsJson, _ := json.Marshal(orgs)
		ls.cache.Put(c, model.MemberOrganization+"::"+memIdStr, string(orgsJson), exp)
	}()
	return &login.LoginResponse{
		Member:           memMsg,
		OrganizationList: orgsMessage,
		TokenList:        tokenList,
	}, nil
}

func (ls *LoginService) TokenVerify(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	token := msg.Token
	if strings.Contains(token, "bearer") {
		token = strings.ReplaceAll(token, "bearer ", "")
	}
	parseToken, err := jwts.ParseToken(token, config.Conf.JwtConfig.AccessSecret, msg.Ip)
	if err != nil {
		zap.L().Error("Login TokenVerify error", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	//数据库查询
	//从缓存中查询，如果缓存没有，则查询失败
	memberJson, err := ls.cache.Get(context.Background(), model.Member+"::"+parseToken)
	if err != nil {
		zap.L().Error("TokenVerify redis cache Get member error", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	if memberJson == "" {
		zap.L().Error("Login TokenVerify cache already expire")
		return nil, errs.GrpcError(model.NoLogin)
	}
	memberById := &member.Member{}
	json.Unmarshal([]byte(memberJson), memberById)
	//后续优化：登录之后，应该把用户信息缓存起来
	//id, _ := strconv.ParseInt(parseToken, 10, 64)
	//memberById, err := ls.memberRepo.FindMemberById(context.Background(), id)

	memMsg := &login.MemberMessage{}
	copier.Copy(memMsg, memberById)
	memMsg.Code, _ = encrypts.EncryptInt64(memberById.Id, model.AESKey)

	orgsJson, err := ls.cache.Get(context.Background(), model.MemberOrganization+"::"+parseToken)
	if err != nil {
		zap.L().Error("TokenVerify redis Get MemberOrganization error", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	if orgsJson == "" {
		zap.L().Error("Login TokenVerify Organization cache already expire")
		return nil, errs.GrpcError(model.NoLogin)
	}
	var orgs []*organization.Organization
	json.Unmarshal([]byte(orgsJson), &orgs)

	//orgs, err := ls.organizationRepo.FindOrganizationByMemId(context.Background(), memberById.Id)
	//if err != nil {
	//	zap.L().Error("TokenVerify FindOrganizationByMemId db fail", zap.Error(err))
	//	return &login.LoginResponse{}, errs.GrpcError(model.DBError)
	//}
	if len(orgs) > 0 { //对第一个组织进行加密
		memMsg.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
	}
	memMsg.CreateTime = tms.FormatByMill(memberById.CreateTime)
	return &login.LoginResponse{
		Member: memMsg,
	}, nil
}

func (l *LoginService) MyOrgList(ctx context.Context, msg *login.UserMessage) (*login.OrgListResponse, error) {
	memId := msg.MemId
	orgs, err := l.organizationRepo.FindOrganizationByMemId(ctx, memId)
	if err != nil {
		zap.L().Error("MyOrgList FindOrganizationByMemId err", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	var orgsMessage []*login.OrganizationMessage
	err = copier.Copy(&orgsMessage, orgs)
	for _, org := range orgsMessage {
		org.Code, _ = encrypts.EncryptInt64(org.Id, model.AESKey)
	}
	return &login.OrgListResponse{OrganizationList: orgsMessage}, nil
}

func (ls *LoginService) FindMemInfoById(ctx context.Context, msg *login.UserMessage) (*login.MemberMessage, error) {
	memberById, err := ls.memberRepo.FindMemberById(context.Background(), msg.MemId)
	if err != nil {
		zap.L().Error("TokenVerify db FindMemberById error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	memMsg := &login.MemberMessage{}
	copier.Copy(memMsg, memberById)
	memMsg.Code, _ = encrypts.EncryptInt64(memberById.Id, model.AESKey)
	orgs, err := ls.organizationRepo.FindOrganizationByMemId(context.Background(), memberById.Id)
	if err != nil {
		zap.L().Error("TokenVerify db FindMember error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if len(orgs) > 0 {
		memMsg.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
	}
	memMsg.CreateTime = tms.FormatByMill(memberById.CreateTime)
	return memMsg, nil
}

func (ls *LoginService) FindMemInfoByIds(ctx context.Context, msg *login.UserMessage) (*login.MemberMessageList, error) {
	memberList, err := ls.memberRepo.FindMemberByIds(context.Background(), msg.MIds)
	if err != nil {
		zap.L().Error("FindMemInfoByIds db memberRepo.FindMemberByIds error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if memberList == nil || len(memberList) <= 0 {
		return &login.MemberMessageList{List: nil}, nil
	}
	mMap := make(map[int64]*member.Member)
	for _, v := range memberList {
		mMap[v.Id] = v
	}
	var memMsgs []*login.MemberMessage
	copier.Copy(&memMsgs, memberList)
	for _, v := range memMsgs {
		m := mMap[v.Id]
		v.CreateTime = tms.FormatByMill(m.CreateTime)
		v.Code = encrypts.EncryptNoErr(v.Id)
	}
	return &login.MemberMessageList{List: memMsgs}, nil
}
