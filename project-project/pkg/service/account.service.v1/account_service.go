package account_service_v1

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/encrypts"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/errs"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/account"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/dao"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/database/tran"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/domain"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/repo"
	"github.com/jinzhu/copier"
)

type AccountService struct {
	account.UnimplementedAccountServiceServer
	cache             repo.Cache
	transaction       tran.Transaction
	accountDomain     *domain.AccountDomain
	projectAuthDomain *domain.ProjectAuthDomain
}

func New() *AccountService {
	return &AccountService{
		cache:             dao.Rc,
		transaction:       dao.NewTransaction(),
		accountDomain:     domain.NewAccountDomain(),
		projectAuthDomain: domain.NewProjectAuthDomain(),
	}
}

func (a *AccountService) Account(c context.Context, msg *account.AccountReqMessage) (*account.AccountResponse, error) {
	accountList, total, err := a.accountDomain.AccountList(
		msg.OrganizationCode,
		msg.MemberId,
		msg.Page,
		msg.PageSize,
		msg.DepartmentCode,
		msg.SearchType)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	authList, err := a.projectAuthDomain.AuthList(encrypts.DecryptNoErr(msg.OrganizationCode))
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var maList []*account.MemberAccount
	copier.Copy(&maList, accountList)
	var prList []*account.ProjectAuth
	copier.Copy(&prList, authList)
	return &account.AccountResponse{
		AccountList: maList,
		AuthList:    prList,
		Total:       total,
	}, nil
}
